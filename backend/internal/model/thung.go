package model

import "time"

// Thung ánh xạ bảng thung (carton lưu kho).
type Thung struct {
	ID          int32     `json:"id"`
	MaThung     string    `json:"ma_thung"`
	SoSerial    *string   `json:"so_serial"`
	DotLuuKho   *int16    `json:"dot_luu_kho"`
	ViTriKho    *string   `json:"vi_tri_kho"`
	NgayNhapKho *Date     `json:"ngay_nhap_kho"`
	GhiChu      *string   `json:"ghi_chu"`
	CreatedAt   time.Time `json:"created_at"`
}

type ThungInput struct {
	MaThung     string  `json:"ma_thung" binding:"required"`
	SoSerial    *string `json:"so_serial"`
	DotLuuKho   *int16  `json:"dot_luu_kho"`
	ViTriKho    *string `json:"vi_tri_kho"`
	NgayNhapKho *Date   `json:"ngay_nhap_kho"`
	GhiChu      *string `json:"ghi_chu"`
}

type ThungFilter struct {
	DotLuuKho *int16
	Search    string // tìm gần đúng trong ma_thung, so_serial
	Page      int
	PageSize  int
}
