package repository

import (
	"fmt"
	"strings"
)

// buildOrderBy dựng mệnh đề ORDER BY từ whitelist cột cho phép — tên cột
// KHÔNG bao giờ lấy trực tiếp từ input người dùng, chỉ tra qua sortColumns.
// sortBy rỗng → dùng defaultOrder. Luôn thêm "id" làm tiêu chí phụ để
// phân trang ổn định.
func buildOrderBy(sortColumns map[string]string, sortBy, sortDir, defaultOrder string) (string, error) {
	if sortBy == "" {
		return defaultOrder, nil
	}

	col, ok := sortColumns[sortBy]
	if !ok {
		allowed := make([]string, 0, len(sortColumns))
		for k := range sortColumns {
			allowed = append(allowed, k)
		}
		return "", fmt.Errorf("%w: sort_by %q không hợp lệ, chỉ chấp nhận: %s",
			ErrInvalidInput, sortBy, strings.Join(allowed, ", "))
	}

	var dir string
	switch strings.ToLower(sortDir) {
	case "", "asc":
		dir = "ASC"
	case "desc":
		dir = "DESC"
	default:
		return "", fmt.Errorf("%w: sort_dir %q không hợp lệ, chỉ chấp nhận asc hoặc desc",
			ErrInvalidInput, sortDir)
	}

	return col + " " + dir + ", id " + dir, nil
}
