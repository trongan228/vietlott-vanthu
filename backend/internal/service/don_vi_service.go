package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type DonViService struct {
	repo *repository.DonViRepository
}

func NewDonViService(repo *repository.DonViRepository) *DonViService {
	return &DonViService{repo: repo}
}

func normalizeDonViInput(in *model.DonViInput) {
	in.TenDonVi = strings.TrimSpace(in.TenDonVi)
	if in.MaDonVi != nil {
		v := strings.TrimSpace(*in.MaDonVi)
		in.MaDonVi = &v
	}
	if in.IsActive == nil {
		t := true
		in.IsActive = &t
	}
}

func (s *DonViService) Create(ctx context.Context, in *model.DonViInput) (*model.DonVi, error) {
	normalizeDonViInput(in)
	return s.repo.Create(ctx, in)
}

func (s *DonViService) GetByID(ctx context.Context, id int32) (*model.DonVi, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DonViService) Update(ctx context.Context, id int32, in *model.DonViInput) (*model.DonVi, error) {
	normalizeDonViInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *DonViService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *DonViService) List(ctx context.Context, f model.DonViFilter) (*model.ListResult[model.DonVi], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.DonVi]{
		Items:      items,
		Pagination: model.Pagination{Page: f.Page, PageSize: f.PageSize, Total: total},
	}, nil
}
