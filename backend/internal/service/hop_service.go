package service

import (
	"context"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type HopService struct {
	repo *repository.HopRepository
}

func NewHopService(repo *repository.HopRepository) *HopService {
	return &HopService{repo: repo}
}

func normalizeHopInput(in *model.HopInput) {
	if in.LoaiHop == "" {
		in.LoaiHop = "thoi_han"
	}
}

func (s *HopService) Create(ctx context.Context, in *model.HopInput) (*model.Hop, error) {
	normalizeHopInput(in)
	return s.repo.Create(ctx, in)
}

func (s *HopService) GetByID(ctx context.Context, id int32) (*model.Hop, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *HopService) Update(ctx context.Context, id int32, in *model.HopInput) (*model.Hop, error) {
	normalizeHopInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *HopService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *HopService) List(ctx context.Context, f model.HopFilter) (*model.ListResult[model.Hop], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.Hop]{
		Items:      items,
		Pagination: model.Pagination{Page: f.Page, PageSize: f.PageSize, Total: total},
	}, nil
}
