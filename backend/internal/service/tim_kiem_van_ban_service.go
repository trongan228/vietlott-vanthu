package service

import (
	"context"
	"fmt"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

const (
	timKiemDefaultLimit = 20
	timKiemMaxLimit     = 50
)

// TimKiemVanBanService tìm kiếm hợp nhất trên cả 3 bảng văn bản.
type TimKiemVanBanService struct {
	repo *repository.TimKiemVanBanRepository
}

func NewTimKiemVanBanService(repo *repository.TimKiemVanBanRepository) *TimKiemVanBanService {
	return &TimKiemVanBanService{repo: repo}
}

func (s *TimKiemVanBanService) TimKiem(ctx context.Context, p model.TimKiemVanBanParams) (*model.TimKiemVanBanKetQua, error) {
	p.Q = strings.TrimSpace(p.Q)
	p.So = strings.TrimSpace(p.So)

	if p.Q == "" && p.So == "" {
		return nil, fmt.Errorf("%w: cần ít nhất một trong hai tham số q (tên/trích yếu) hoặc so (số văn bản)",
			repository.ErrInvalidInput)
	}
	if p.Limit < 1 {
		p.Limit = timKiemDefaultLimit
	}
	if p.Limit > timKiemMaxLimit {
		p.Limit = timKiemMaxLimit
	}

	den, err := s.repo.TimVanBanDen(ctx, p)
	if err != nil {
		return nil, err
	}
	diKhongTenLoai, err := s.repo.TimVanBanDiKhongTenLoai(ctx, p)
	if err != nil {
		return nil, err
	}
	diCoTenLoai, err := s.repo.TimVanBanDiCoTenLoai(ctx, p)
	if err != nil {
		return nil, err
	}

	// Gom ho_so_id của cả 3 danh sách, tra vị trí lưu kho một lần
	// rồi gắn vi_tri (hồ sơ/hộp/thùng) vào từng kết quả.
	idSet := make(map[int32]struct{})
	for _, m := range den {
		if m.HoSoID != nil {
			idSet[*m.HoSoID] = struct{}{}
		}
	}
	for _, m := range diKhongTenLoai {
		if m.HoSoID != nil {
			idSet[*m.HoSoID] = struct{}{}
		}
	}
	for _, m := range diCoTenLoai {
		if m.HoSoID != nil {
			idSet[*m.HoSoID] = struct{}{}
		}
	}
	hoSoIDs := make([]int32, 0, len(idSet))
	for id := range idSet {
		hoSoIDs = append(hoSoIDs, id)
	}
	viTri, err := s.repo.LayViTriLuuKho(ctx, hoSoIDs)
	if err != nil {
		return nil, err
	}

	ketQua := &model.TimKiemVanBanKetQua{
		VanBanDen:            make([]model.VanBanDenTimKiem, len(den)),
		VanBanDiKhongTenLoai: make([]model.VanBanDiKhongTenLoaiTimKiem, len(diKhongTenLoai)),
		VanBanDiCoTenLoai:    make([]model.VanBanDiCoTenLoaiTimKiem, len(diCoTenLoai)),
	}
	for i, m := range den {
		ketQua.VanBanDen[i] = model.VanBanDenTimKiem{VanBanDen: m, ViTri: timViTri(viTri, m.HoSoID)}
	}
	for i, m := range diKhongTenLoai {
		ketQua.VanBanDiKhongTenLoai[i] = model.VanBanDiKhongTenLoaiTimKiem{VanBanDiKhongTenLoai: m, ViTri: timViTri(viTri, m.HoSoID)}
	}
	for i, m := range diCoTenLoai {
		ketQua.VanBanDiCoTenLoai[i] = model.VanBanDiCoTenLoaiTimKiem{VanBanDiCoTenLoai: m, ViTri: timViTri(viTri, m.HoSoID)}
	}
	return ketQua, nil
}

// timViTri tra vị trí lưu kho theo ho_so_id; trả nil nếu văn bản chưa gán hồ sơ.
func timViTri(viTri map[int32]model.ViTriLuuKho, hoSoID *int32) *model.ViTriLuuKho {
	if hoSoID == nil {
		return nil
	}
	if v, ok := viTri[*hoSoID]; ok {
		return &v
	}
	return nil
}
