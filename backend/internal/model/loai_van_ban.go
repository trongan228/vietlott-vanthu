package model

// LoaiVanBan ánh xạ bảng loai_van_ban (danh mục loại văn bản).
type LoaiVanBan struct {
	ID                int32   `json:"id"`
	MaLoai            string  `json:"ma_loai"`
	TenLoai           string  `json:"ten_loai"`
	ThoiHanBaoQuanNam *int32  `json:"thoi_han_bao_quan_nam"`
	GhiChu            *string `json:"ghi_chu"`
}

type LoaiVanBanInput struct {
	MaLoai            string  `json:"ma_loai" binding:"required"`
	TenLoai           string  `json:"ten_loai" binding:"required"`
	ThoiHanBaoQuanNam *int32  `json:"thoi_han_bao_quan_nam"`
	GhiChu            *string `json:"ghi_chu"`
}

type LoaiVanBanFilter struct {
	Search   string // tìm gần đúng trong ma_loai, ten_loai
	Page     int
	PageSize int
}
