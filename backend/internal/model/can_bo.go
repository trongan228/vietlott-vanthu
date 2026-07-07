package model

import "time"

// CanBo ánh xạ bảng can_bo (danh mục cán bộ).
type CanBo struct {
	ID          int32     `json:"id"`
	HoTen       string    `json:"ho_ten"`
	ChucDanh    *string   `json:"chuc_danh"`
	DonViID     *int32    `json:"don_vi_id"`
	SoDienThoai *string   `json:"so_dien_thoai"`
	Email       *string   `json:"email"`
	IsVanThu    bool      `json:"is_van_thu"`
	IsActive    bool      `json:"is_active"`
	GhiChu      *string   `json:"ghi_chu"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CanBoInput struct {
	HoTen       string  `json:"ho_ten" binding:"required"`
	ChucDanh    *string `json:"chuc_danh"`
	DonViID     *int32  `json:"don_vi_id"`
	SoDienThoai *string `json:"so_dien_thoai"`
	Email       *string `json:"email" binding:"omitempty,email"`
	IsVanThu    *bool   `json:"is_van_thu"` // nil = mặc định FALSE
	IsActive    *bool   `json:"is_active"`  // nil = mặc định TRUE
	GhiChu      *string `json:"ghi_chu"`
}

type CanBoFilter struct {
	DonViID  *int32
	IsVanThu *bool
	IsActive *bool
	Search   string // tìm gần đúng trong ho_ten
	Page     int
	PageSize int
}
