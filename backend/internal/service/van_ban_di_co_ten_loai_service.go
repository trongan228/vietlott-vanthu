package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type VanBanDiCoTenLoaiService struct {
	repo *repository.VanBanDiCoTenLoaiRepository
}

func NewVanBanDiCoTenLoaiService(repo *repository.VanBanDiCoTenLoaiRepository) *VanBanDiCoTenLoaiService {
	return &VanBanDiCoTenLoaiService{repo: repo}
}

func normalizeVanBanDiCoTenLoaiInput(in *model.VanBanDiCoTenLoaiInput) {
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

func (s *VanBanDiCoTenLoaiService) Create(ctx context.Context, in *model.VanBanDiCoTenLoaiInput) (*model.VanBanDiCoTenLoai, error) {
	normalizeVanBanDiCoTenLoaiInput(in)
	return s.repo.Create(ctx, in)
}

func (s *VanBanDiCoTenLoaiService) GetByID(ctx context.Context, id int64) (*model.VanBanDiCoTenLoai, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *VanBanDiCoTenLoaiService) Update(ctx context.Context, id int64, in *model.VanBanDiCoTenLoaiInput) (*model.VanBanDiCoTenLoai, error) {
	normalizeVanBanDiCoTenLoaiInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *VanBanDiCoTenLoaiService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *VanBanDiCoTenLoaiService) List(ctx context.Context, f model.VanBanDiCoTenLoaiFilter) (*model.ListResult[model.VanBanDiCoTenLoai], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.VanBanDiCoTenLoai]{
		Items: items,
		Pagination: model.Pagination{
			Page:     f.Page,
			PageSize: f.PageSize,
			Total:    total,
		},
	}, nil
}
