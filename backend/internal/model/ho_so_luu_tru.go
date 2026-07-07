package model

import "time"

// HoSoLuuTru ánh xạ bảng ho_so_luu_tru (tập lưu văn bản).
type HoSoLuuTru struct {
	ID                int32     `json:"id"`
	SoHoSo            int32     `json:"so_ho_so"`
	HopID             int32     `json:"hop_id"`
	TieuDe            string    `json:"tieu_de"`
	LoaiTap           string    `json:"loai_tap"`
	Nam               *int16    `json:"nam"`
	TuNgay            *Date     `json:"tu_ngay"`
	DenNgay           *Date     `json:"den_ngay"`
	TuSo              *int32    `json:"tu_so"`
	DenSo             *int32    `json:"den_so"`
	SoTo              *int32    `json:"so_to"`
	ThoiHanBaoQuanNam *int32    `json:"thoi_han_bao_quan_nam"`
	VinhVien          bool      `json:"vinh_vien"`
	GhiChu            *string   `json:"ghi_chu"`
	CreatedAt         time.Time `json:"created_at"`
}

type HoSoLuuTruInput struct {
	SoHoSo            int32   `json:"so_ho_so" binding:"required"`
	HopID             int32   `json:"hop_id" binding:"required"`
	TieuDe            string  `json:"tieu_de" binding:"required"`
	LoaiTap           string  `json:"loai_tap" binding:"required,oneof=CV_DEN CV_DI_CO_TEN_LOAI CV_DI_KHONG_TEN_LOAI CV_TO_BAN_HD CV_DOAN_THE CV_CHI_NHANH CV_NOI_BO HDLD HD_LAY_SO KHAC"`
	Nam               *int16  `json:"nam"`
	TuNgay            *Date   `json:"tu_ngay"`
	DenNgay           *Date   `json:"den_ngay"`
	TuSo              *int32  `json:"tu_so"`
	DenSo             *int32  `json:"den_so"`
	SoTo              *int32  `json:"so_to"`
	ThoiHanBaoQuanNam *int32  `json:"thoi_han_bao_quan_nam"`
	VinhVien          *bool   `json:"vinh_vien"` // nil = mặc định FALSE
	GhiChu            *string `json:"ghi_chu"`
}

type HoSoLuuTruFilter struct {
	HopID    *int32
	LoaiTap  *string
	Nam      *int16
	VinhVien *bool
	Search   string // tìm gần đúng trong tieu_de
	Page     int
	PageSize int
}
