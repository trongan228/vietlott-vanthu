package model

// ViTriVanBanDen ánh xạ view v_vi_tri_van_ban_den:
// văn bản đến → hồ sơ → hộp → thùng.
type ViTriVanBanDen struct {
	VanBanID   int64   `json:"van_ban_id"`
	Nam        int16   `json:"nam"`
	SoDen      int32   `json:"so_den"`
	NgayDen    Date    `json:"ngay_den"`
	TrichYeu   string  `json:"trich_yeu"`
	SoHoSo     int32   `json:"so_ho_so"`
	HoSoTieuDe string  `json:"ho_so_tieu_de"`
	SoHop      int32   `json:"so_hop"`
	MaThung    *string `json:"ma_thung"` // NULL nếu hộp chưa xếp vào thùng
	SoSerial   *string `json:"so_serial"`
}

// ViTriVanBanDi ánh xạ 2 view v_vi_tri_van_ban_di_* (cùng cấu trúc cột).
type ViTriVanBanDi struct {
	VanBanID   int64   `json:"van_ban_id"`
	Nam        int16   `json:"nam"`
	SoKyHieu   string  `json:"so_ky_hieu"`
	NgayVanBan Date    `json:"ngay_van_ban"`
	TrichYeu   string  `json:"trich_yeu"`
	SoHoSo     int32   `json:"so_ho_so"`
	HoSoTieuDe string  `json:"ho_so_tieu_de"`
	SoHop      int32   `json:"so_hop"`
	MaThung    *string `json:"ma_thung"`
	SoSerial   *string `json:"so_serial"`
}
