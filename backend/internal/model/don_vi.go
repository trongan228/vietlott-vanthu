package model

import "time"

// DonVi ánh xạ bảng don_vi (danh mục đơn vị).
type DonVi struct {
	ID        int32     `json:"id"`
	MaDonVi   *string   `json:"ma_don_vi"`
	TenDonVi  string    `json:"ten_don_vi"`
	LoaiDonVi string    `json:"loai_don_vi"`
	DiaChi    *string   `json:"dia_chi"`
	GhiChu    *string   `json:"ghi_chu"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DonViInput struct {
	MaDonVi   *string `json:"ma_don_vi"`
	TenDonVi  string  `json:"ten_don_vi" binding:"required"`
	LoaiDonVi string  `json:"loai_don_vi" binding:"required,oneof=noi_bo chi_nhanh doi_tac to_chuc"`
	DiaChi    *string `json:"dia_chi"`
	GhiChu    *string `json:"ghi_chu"`
	IsActive  *bool   `json:"is_active"` // nil = giữ mặc định TRUE
}

type DonViFilter struct {
	LoaiDonVi *string
	IsActive  *bool
	Search    string // tìm gần đúng trong ten_don_vi, ma_don_vi
	Page      int
	PageSize  int
}
