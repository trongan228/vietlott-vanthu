package model

import "time"

// Hop ánh xạ bảng hop. Hộp "thoi_han" và "vinh_vien" là hai dãy số riêng,
// khóa duy nhất là (so_hop, loai_hop).
type Hop struct {
	ID        int32     `json:"id"`
	SoHop     int32     `json:"so_hop"`
	LoaiHop   string    `json:"loai_hop"`
	ThungID   *int32    `json:"thung_id"`
	GhiChu    *string   `json:"ghi_chu"`
	CreatedAt time.Time `json:"created_at"`
}

type HopInput struct {
	SoHop   int32   `json:"so_hop" binding:"required"`
	LoaiHop string  `json:"loai_hop" binding:"omitempty,oneof=thoi_han vinh_vien"` // rỗng = thoi_han
	ThungID *int32  `json:"thung_id"`
	GhiChu  *string `json:"ghi_chu"`
}

type HopFilter struct {
	SoHop    *int32
	LoaiHop  *string
	ThungID  *int32
	Page     int
	PageSize int
}
