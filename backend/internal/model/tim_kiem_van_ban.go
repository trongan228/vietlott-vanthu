package model

// TimKiemVanBanParams là tham số tìm kiếm hợp nhất trên cả 3 bảng văn bản.
// Phải có ít nhất một trong Q hoặc So.
type TimKiemVanBanParams struct {
	Q     string // tìm gần đúng theo tên/trích yếu, số văn bản (so_ky_hieu); với văn bản đến khớp thêm số đến (so_den)
	So    string // tìm gần đúng chỉ theo số văn bản (so_ky_hieu; với văn bản đến khớp thêm so_den)
	Nam   *int16
	Limit int // số kết quả tối đa mỗi bảng
}

// ViTriLuuKho là vị trí lưu kho của văn bản theo hồ sơ đã gán
// (ho_so_id): văn bản → hồ sơ → hộp → thùng.
type ViTriLuuKho struct {
	HoSoID     int32   `json:"ho_so_id"`
	SoHoSo     int32   `json:"so_ho_so"`
	HoSoTieuDe string  `json:"ho_so_tieu_de"`
	SoHop      int32   `json:"so_hop"`
	LoaiHop    string  `json:"loai_hop"` // thoi_han | vinh_vien
	MaThung    *string `json:"ma_thung"` // NULL nếu hộp chưa xếp vào thùng
	SoSerial   *string `json:"so_serial"`
}

// Các item kết quả tìm kiếm: văn bản kèm vị trí lưu kho (nếu đã gán hồ sơ).
type VanBanDenTimKiem struct {
	VanBanDen
	ViTri *ViTriLuuKho `json:"vi_tri"`
}

type VanBanDiKhongTenLoaiTimKiem struct {
	VanBanDiKhongTenLoai
	ViTri *ViTriLuuKho `json:"vi_tri"`
}

type VanBanDiCoTenLoaiTimKiem struct {
	VanBanDiCoTenLoai
	ViTri *ViTriLuuKho `json:"vi_tri"`
}

// TimKiemVanBanKetQua gom kết quả từ 3 bảng văn bản.
type TimKiemVanBanKetQua struct {
	VanBanDen            []VanBanDenTimKiem            `json:"van_ban_den"`
	VanBanDiKhongTenLoai []VanBanDiKhongTenLoaiTimKiem `json:"van_ban_di_khong_ten_loai"`
	VanBanDiCoTenLoai    []VanBanDiCoTenLoaiTimKiem    `json:"van_ban_di_co_ten_loai"`
}
