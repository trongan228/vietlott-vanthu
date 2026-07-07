package model

import "time"

// VanBanDen ánh xạ bảng van_ban_den (migrations/00_schema_chinh.sql).
type VanBanDen struct {
	ID            int64     `json:"id"`
	Nam           int16     `json:"nam"`
	SoDen         int32     `json:"so_den"`
	NgayDen       Date      `json:"ngay_den"`
	NoiGuiID      *int32    `json:"noi_gui_id"`
	NoiGuiText    string    `json:"noi_gui_text"`
	SoKyHieu      *string   `json:"so_ky_hieu"`
	NgayVanBan    *Date     `json:"ngay_van_ban"`
	TrichYeu      string    `json:"trich_yeu"`
	DonViXuLyID   *int32    `json:"don_vi_xu_ly_id"`
	DonViNhanText *string   `json:"don_vi_nhan_text"`
	KyNhan        *string   `json:"ky_nhan"`
	GhiChu        *string   `json:"ghi_chu"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// VanBanDenInput là payload cho POST/PUT.
type VanBanDenInput struct {
	Nam           int16   `json:"nam" binding:"required"`
	SoDen         int32   `json:"so_den" binding:"required"`
	NgayDen       Date    `json:"ngay_den" binding:"required"`
	NoiGuiID      *int32  `json:"noi_gui_id"`
	NoiGuiText    string  `json:"noi_gui_text" binding:"required"`
	SoKyHieu      *string `json:"so_ky_hieu"`
	NgayVanBan    *Date   `json:"ngay_van_ban"`
	TrichYeu      string  `json:"trich_yeu" binding:"required"`
	DonViXuLyID   *int32  `json:"don_vi_xu_ly_id"`
	DonViNhanText *string `json:"don_vi_nhan_text"`
	KyNhan        *string `json:"ky_nhan"`
	GhiChu        *string `json:"ghi_chu"`
}

// VanBanDenFilter dùng cho GET danh sách (List).
type VanBanDenFilter struct {
	Nam         *int16
	SoDen       *int32
	SoKyHieu    *string // tìm gần đúng trong so_ky_hieu
	NoiGuiID    *int32  // lọc theo đơn vị ban hành (nơi gửi)
	DonViXuLyID *int32  // lọc theo đơn vị xử lý
	Search      string  // tìm gần đúng trong trich_yeu
	SortBy      string  // so_den | ngay_den | ngay_van_ban | nam | created_at
	SortDir     string  // asc | desc
	Page        int
	PageSize    int
}
