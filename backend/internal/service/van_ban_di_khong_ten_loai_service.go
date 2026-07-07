package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type VanBanDiKhongTenLoaiService struct {
	repo *repository.VanBanDiKhongTenLoaiRepository
}

func NewVanBanDiKhongTenLoaiService(repo *repository.VanBanDiKhongTenLoaiRepository) *VanBanDiKhongTenLoaiService {
	return &VanBanDiKhongTenLoaiService{repo: repo}
}

func normalizeVanBanDiKhongTenLoaiInput(in *model.VanBanDiKhongTenLoaiInput) {
	in.SoKyHieu = strings.TrimSpace(in.SoKyHieu)
	in.TrichYeu = strings.TrimSpace(in.TrichYeu)
	if in.NguoiKyText != nil {
		v := strings.TrimSpace(*in.NguoiKyText)
		in.NguoiKyText = &v
	}
	if in.NoiNhanText != nil {
		v := strings.TrimSpace(*in.NoiNhanText)
		in.NoiNhanText = &v
	}
}

func (s *VanBanDiKhongTenLoaiService) Create(ctx context.Context, in *model.VanBanDiKhongTenLoaiInput) (*model.VanBanDiKhongTenLoai, error) {
	normalizeVanBanDiKhongTenLoaiInput(in)
	return s.repo.Create(ctx, in)
}

func (s *VanBanDiKhongTenLoaiService) GetByID(ctx context.Context, id int64) (*model.VanBanDiKhongTenLoai, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *VanBanDiKhongTenLoaiService) Update(ctx context.Context, id int64, in *model.VanBanDiKhongTenLoaiInput) (*model.VanBanDiKhongTenLoai, error) {
	normalizeVanBanDiKhongTenLoaiInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *VanBanDiKhongTenLoaiService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *VanBanDiKhongTenLoaiService) List(ctx context.Context, f model.VanBanDiKhongTenLoaiFilter) (*model.ListResult[model.VanBanDiKhongTenLoai], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.VanBanDiKhongTenLoai]{
		Items: items,
		Pagination: model.Pagination{
			Page:     f.Page,
			PageSize: f.PageSize,
			Total:    total,
		},
	}, nil
}
