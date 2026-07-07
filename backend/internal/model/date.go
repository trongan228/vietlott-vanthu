package model

import (
	"fmt"
	"strings"
	"time"
)

// DateLayout là định dạng ngày dùng cho JSON (không có giờ), khớp với
// các cột kiểu DATE trong schema (ngay_den, ngay_van_ban...).
const DateLayout = "2006-01-02"

// Date bọc time.Time để (un)marshal JSON theo định dạng "YYYY-MM-DD"
// thay vì RFC3339 mặc định của time.Time.
type Date struct {
	time.Time
}

func DateFromTime(t time.Time) Date {
	return Date{Time: t}
}

func DatePtrFromTimePtr(t *time.Time) *Date {
	if t == nil {
		return nil
	}
	d := Date{Time: *t}
	return &d
}

// TimePtr trả về *time.Time tương ứng, an toàn khi gọi trên con trỏ nil.
func (d *Date) TimePtr() *time.Time {
	if d == nil {
		return nil
	}
	t := d.Time
	return &t
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(DateLayout) + `"`), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "null" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return fmt.Errorf("ngày không hợp lệ %q, cần định dạng YYYY-MM-DD", s)
	}
	d.Time = t
	return nil
}
