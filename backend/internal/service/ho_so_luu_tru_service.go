package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type HoSoLuuTruService struct {
	repo *repository.HoSoLuuTruRepository
}

func NewHoSoLuuTruService(repo *repository.HoSoLuuTruRepository) *HoSoLuuTruService {
	return &HoSoLuuTruService{repo: repo}
}

func normalizeHoSoLuuTruInput(in *model.HoSoLuuTruInput) {
	in.TieuDe = strings.TrimSpace(in.TieuDe)
	if in.VinhVien == nil {
		f := false
		in.VinhVien = &f
	}
}

func (s *HoSoLuuTruService) Create(ctx context.Context, in *model.HoSoLuuTruInput) (*model.HoSoLuuTru, error) {
	normalizeHoSoLuuTruInput(in)
	return s.repo.Create(ctx, in)
}

func (s *HoSoLuuTruService) GetByID(ctx context.Context, id int32) (*model.HoSoLuuTru, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *HoSoLuuTruService) Update(ctx context.Context, id int32, in *model.HoSoLuuTruInput) (*model.HoSoLuuTru, error) {
	normalizeHoSoLuuTruInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *HoSoLuuTruService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *HoSoLuuTruService) List(ctx context.Context, f model.HoSoLuuTruFilter) (*model.ListResult[model.HoSoLuuTru], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.HoSoLuuTru]{
		Items:      items,
		Pagination: model.Pagination{Page: f.Page, PageSize: f.PageSize, Total: total},
	}, nil
}
