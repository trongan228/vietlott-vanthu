package repository

import "errors"

// ErrNotFound trả về khi không tìm thấy bản ghi theo id.
var ErrNotFound = errors.New("record not found")

// ErrInvalidInput trả về khi tham số lọc/sắp xếp không hợp lệ (map sang 400).
var ErrInvalidInput = errors.New("invalid input")
