package model

// TimKiemVanBanParams là tham số tìm kiếm hợp nhất trên cả 3 bảng văn bản.
// Phải có ít nhất một trong Q hoặc So.
type TimKiemVanBanParams struct {
	Q     string // tìm gần đúng trong trich_yeu (tên/nội dung văn bản)
	So    string // tìm gần đúng theo số văn bản (so_ky_hieu; với văn bản đến khớp thêm so_den)
	Nam   *int16
	Limit int // số kết quả tối đa mỗi bảng
}

// TimKiemVanBanKetQua gom kết quả từ 3 bảng văn bản.
type TimKiemVanBanKetQua struct {
	VanBanDen            []VanBanDen            `json:"van_ban_den"`
	VanBanDiKhongTenLoai []VanBanDiKhongTenLoai `json:"van_ban_di_khong_ten_loai"`
	VanBanDiCoTenLoai    []VanBanDiCoTenLoai    `json:"van_ban_di_co_ten_loai"`
}
