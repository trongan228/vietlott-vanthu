package service

import (
	"context"
	"fmt"
	"strings"

	"vanthu-backend/internal/model"
	"vanthu-backend/internal/repository"
)

const (
	timKiemDefaultLimit = 20
	timKiemMaxLimit     = 50
)

// TimKiemVanBanService tìm kiếm hợp nhất trên cả 3 bảng văn bản.
type TimKiemVanBanService struct {
	repo *repository.TimKiemVanBanRepository
}

func NewTimKiemVanBanService(repo *repository.TimKiemVanBanRepository) *TimKiemVanBanService {
	return &TimKiemVanBanService{repo: repo}
}

func (s *TimKiemVanBanService) TimKiem(ctx context.Context, p model.TimKiemVanBanParams) (*model.TimKiemVanBanKetQua, error) {
	p.Q = strings.TrimSpace(p.Q)
	p.So = strings.TrimSpace(p.So)

	if p.Q == "" && p.So == "" {
		return nil, fmt.Errorf("%w: cần ít nhất một trong hai tham số q (tên/trích yếu) hoặc so (số văn bản)",
			repository.ErrInvalidInput)
	}
	if p.Limit < 1 {
		p.Limit = timKiemDefaultLimit
	}
	if p.Limit > timKiemMaxLimit {
		p.Limit = timKiemMaxLimit
	}

	den, err := s.repo.TimVanBanDen(ctx, p)
	if err != nil {
		return nil, err
	}
	diKhongTenLoai, err := s.repo.TimVanBanDiKhongTenLoai(ctx, p)
	if err != nil {
		return nil, err
	}
	diCoTenLoai, err := s.repo.TimVanBanDiCoTenLoai(ctx, p)
	if err != nil {
		return nil, err
	}

	return &model.TimKiemVanBanKetQua{
		VanBanDen:            den,
		VanBanDiKhongTenLoai: diKhongTenLoai,
		VanBanDiCoTenLoai:    diCoTenLoai,
	}, nil
}
