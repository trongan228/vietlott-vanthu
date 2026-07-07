package service

import (
	"context"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

// ViTriVanBanService tra cứu vị trí lưu kho của văn bản (chỉ đọc).
type ViTriVanBanService struct {
	repo *repository.ViTriVanBanRepository
}

func NewViTriVanBanService(repo *repository.ViTriVanBanRepository) *ViTriVanBanService {
	return &ViTriVanBanService{repo: repo}
}

func (s *ViTriVanBanService) TraCuuVanBanDen(ctx context.Context, nam *int16, soDen *int32) ([]model.ViTriVanBanDen, error) {
	return s.repo.TraCuuVanBanDen(ctx, nam, soDen)
}

func (s *ViTriVanBanService) TraCuuVanBanDiCoTenLoai(ctx context.Context, nam *int16, soKyHieu *string) ([]model.ViTriVanBanDi, error) {
	return s.repo.TraCuuVanBanDiCoTenLoai(ctx, nam, soKyHieu)
}

func (s *ViTriVanBanService) TraCuuVanBanDiKhongTenLoai(ctx context.Context, nam *int16, soKyHieu *string) ([]model.ViTriVanBanDi, error) {
	return s.repo.TraCuuVanBanDiKhongTenLoai(ctx, nam, soKyHieu)
}
