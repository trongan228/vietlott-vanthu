package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type CanBoService struct {
	repo *repository.CanBoRepository
}

func NewCanBoService(repo *repository.CanBoRepository) *CanBoService {
	return &CanBoService{repo: repo}
}

func normalizeCanBoInput(in *model.CanBoInput) {
	in.HoTen = strings.TrimSpace(in.HoTen)
	if in.ChucDanh != nil {
		v := strings.TrimSpace(*in.ChucDanh)
		in.ChucDanh = &v
	}
	if in.Email != nil {
		v := strings.TrimSpace(*in.Email)
		in.Email = &v
	}
	if in.IsVanThu == nil {
		f := false
		in.IsVanThu = &f
	}
	if in.IsActive == nil {
		t := true
		in.IsActive = &t
	}
}

func (s *CanBoService) Create(ctx context.Context, in *model.CanBoInput) (*model.CanBo, error) {
	normalizeCanBoInput(in)
	return s.repo.Create(ctx, in)
}

func (s *CanBoService) GetByID(ctx context.Context, id int32) (*model.CanBo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CanBoService) Update(ctx context.Context, id int32, in *model.CanBoInput) (*model.CanBo, error) {
	normalizeCanBoInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *CanBoService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *CanBoService) List(ctx context.Context, f model.CanBoFilter) (*model.ListResult[model.CanBo], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.CanBo]{
		Items:      items,
		Pagination: model.Pagination{Page: f.Page, PageSize: f.PageSize, Total: total},
	}, nil
}
