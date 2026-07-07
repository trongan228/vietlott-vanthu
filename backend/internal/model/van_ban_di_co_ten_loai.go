package model

import "time"

// VanBanDiCoTenLoai ánh xạ bảng van_ban_di_co_ten_loai.
type VanBanDiCoTenLoai struct {
	ID           int64     `json:"id"`
	Nam          int16     `json:"nam"`
	LoaiVanBanID *int32    `json:"loai_van_ban_id"`
	SoKyHieu     string    `json:"so_ky_hieu"`
	NgayVanBan   Date      `json:"ngay_van_ban"`
	TrichYeu     string    `json:"trich_yeu"`
	NguoiKyID    *int32    `json:"nguoi_ky_id"`
	NguoiKyText  *string   `json:"nguoi_ky_text"`
	NoiNhanText  *string   `json:"noi_nhan_text"`
	SoLuongBan   *int32    `json:"so_luong_ban"`
	GhiChu       *string   `json:"ghi_chu"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type VanBanDiCoTenLoaiInput struct {
	Nam          int16   `json:"nam" binding:"required"`
	LoaiVanBanID *int32  `json:"loai_van_ban_id"`
	SoKyHieu     string  `json:"so_ky_hieu" binding:"required"`
	NgayVanBan   Date    `json:"ngay_van_ban" binding:"required"`
	TrichYeu     string  `json:"trich_yeu" binding:"required"`
	NguoiKyID    *int32  `json:"nguoi_ky_id"`
	NguoiKyText  *string `json:"nguoi_ky_text"`
	NoiNhanText  *string `json:"noi_nhan_text"`
	SoLuongBan   *int32  `json:"so_luong_ban"`
	GhiChu       *string `json:"ghi_chu"`
}

type VanBanDiCoTenLoaiFilter struct {
	Nam          *int16
	LoaiVanBanID *int32  // lọc theo loại văn bản
	SoKyHieu     *string // tìm gần đúng trong so_ky_hieu
	NguoiKyID    *int32  // lọc theo người ký
	Search       string  // tìm gần đúng trong trich_yeu
	SortBy       string  // so_ky_hieu | ngay_van_ban | nam | created_at
	SortDir      string  // asc | desc
	Page         int
	PageSize     int
}
