package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type VanBanDenService struct {
	repo *repository.VanBanDenRepository
}

func NewVanBanDenService(repo *repository.VanBanDenRepository) *VanBanDenService {
	return &VanBanDenService{repo: repo}
}

func normalizeVanBanDenInput(in *model.VanBanDenInput) {
	in.NoiGuiText = strings.TrimSpace(in.NoiGuiText)
	in.TrichYeu = strings.TrimSpace(in.TrichYeu)
	if in.SoKyHieu != nil {
		v := strings.TrimSpace(*in.SoKyHieu)
		in.SoKyHieu = &v
	}
	if in.DonViNhanText != nil {
		v := strings.TrimSpace(*in.DonViNhanText)
		in.DonViNhanText = &v
	}
	if in.KyNhan != nil {
		v := strings.TrimSpace(*in.KyNhan)
		in.KyNhan = &v
	}
}

func (s *VanBanDenService) Create(ctx context.Context, in *model.VanBanDenInput) (*model.VanBanDen, error) {
	normalizeVanBanDenInput(in)
	return s.repo.Create(ctx, in)
}

func (s *VanBanDenService) GetByID(ctx context.Context, id int64) (*model.VanBanDen, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *VanBanDenService) Update(ctx context.Context, id int64, in *model.VanBanDenInput) (*model.VanBanDen, error) {
	normalizeVanBanDenInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *VanBanDenService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *VanBanDenService) List(ctx context.Context, f model.VanBanDenFilter) (*model.ListResult[model.VanBanDen], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.VanBanDen]{
		Items: items,
		Pagination: model.Pagination{
			Page:     f.Page,
			PageSize: f.PageSize,
			Total:    total,
		},
	}, nil
}
