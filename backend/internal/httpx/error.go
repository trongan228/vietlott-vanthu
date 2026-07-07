package httpx

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"

	"vanthu-backend/internal/repository"
)

// WriteError ánh xạ lỗi từ service/repository sang HTTP response phù hợp.
func WriteError(c *gin.Context, err error) {
	if errors.Is(err, repository.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "không tìm thấy bản ghi"})
		return
	}

	if errors.Is(err, repository.ErrInvalidInput) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			c.JSON(http.StatusConflict, gin.H{
				"error":  "dữ liệu đã tồn tại (trùng năm + số/ký hiệu)",
				"detail": pgErr.Detail,
			})
			return
		case "23503": // foreign_key_violation
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "tham chiếu không hợp lệ (đơn vị/cán bộ/loại văn bản không tồn tại)",
				"detail": pgErr.Detail,
			})
			return
		case "23514": // check_violation
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "dữ liệu không hợp lệ",
				"detail": pgErr.Detail,
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "lỗi hệ thống", "detail": err.Error()})
}
