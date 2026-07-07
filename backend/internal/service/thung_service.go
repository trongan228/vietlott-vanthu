package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type ThungService struct {
	repo *repository.ThungRepository
}

func NewThungService(repo *repository.ThungRepository) *ThungService {
	return &ThungService{repo: repo}
}

func normalizeThungInput(in *model.ThungInput) {
	in.MaThung = strings.TrimSpace(in.MaThung)
	if in.SoSerial != nil {
		v := strings.TrimSpace(*in.SoSerial)
		in.SoSerial = &v
	}
}

func (s *ThungService) Create(ctx context.Context, in *model.ThungInput) (*model.Thung, error) {
	normalizeThungInput(in)
	return s.repo.Create(ctx, in)
}

func (s *ThungService) GetByID(ctx context.Context, id int32) (*model.Thung, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ThungService) Update(ctx context.Context, id int32, in *model.ThungInput) (*model.Thung, error) {
	normalizeThungInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *ThungService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *ThungService) List(ctx context.Context, f model.ThungFilter) (*model.ListResult[model.Thung], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.Thung]{
		Items:      items,
		Pagination: model.Pagination{Page: f.Page, PageSize: f.PageSize, Total: total},
	}, nil
}
