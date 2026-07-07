package service

import (
	"context"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

type LoaiVanBanService struct {
	repo *repository.LoaiVanBanRepository
}

func NewLoaiVanBanService(repo *repository.LoaiVanBanRepository) *LoaiVanBanService {
	return &LoaiVanBanService{repo: repo}
}

func normalizeLoaiVanBanInput(in *model.LoaiVanBanInput) {
	in.MaLoai = strings.TrimSpace(in.MaLoai)
	in.TenLoai = strings.TrimSpace(in.TenLoai)
}

func (s *LoaiVanBanService) Create(ctx context.Context, in *model.LoaiVanBanInput) (*model.LoaiVanBan, error) {
	normalizeLoaiVanBanInput(in)
	return s.repo.Create(ctx, in)
}

func (s *LoaiVanBanService) GetByID(ctx context.Context, id int32) (*model.LoaiVanBan, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LoaiVanBanService) Update(ctx context.Context, id int32, in *model.LoaiVanBanInput) (*model.LoaiVanBan, error) {
	normalizeLoaiVanBanInput(in)
	return s.repo.Update(ctx, id, in)
}

func (s *LoaiVanBanService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *LoaiVanBanService) List(ctx context.Context, f model.LoaiVanBanFilter) (*model.ListResult[model.LoaiVanBan], error) {
	f.Page, f.PageSize = model.NormalizePaging(f.Page, f.PageSize)

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.ListResult[model.LoaiVanBan]{
		Items:      items,
		Pagination: model.Pagination{Page: f.Page, PageSize: f.PageSize, Total: total},
	}, nil
}
