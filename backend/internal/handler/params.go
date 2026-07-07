package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// parseIDParam32 đọc path param :id kiểu int32; trả về ok=false nếu đã
// ghi response lỗi 400.
func parseIDParam32(c *gin.Context) (int32, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return 0, false
	}
	return int32(id), true
}

// queryInt16 đọc query param kiểu *int16; trả về ok=false nếu đã ghi lỗi 400.
func queryInt16(c *gin.Context, name string) (*int16, bool) {
	v := c.Query(name)
	if v == "" {
		return nil, true
	}
	n, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tham số " + name + " không hợp lệ"})
		return nil, false
	}
	nn := int16(n)
	return &nn, true
}

// queryInt32 đọc query param kiểu *int32; trả về ok=false nếu đã ghi lỗi 400.
func queryInt32(c *gin.Context, name string) (*int32, bool) {
	v := c.Query(name)
	if v == "" {
		return nil, true
	}
	n, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tham số " + name + " không hợp lệ"})
		return nil, false
	}
	nn := int32(n)
	return &nn, true
}

// queryBool đọc query param kiểu *bool (true/false); trả về ok=false nếu đã ghi lỗi 400.
func queryBool(c *gin.Context, name string) (*bool, bool) {
	v := c.Query(name)
	if v == "" {
		return nil, true
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tham số " + name + " không hợp lệ (true/false)"})
		return nil, false
	}
	return &b, true
}

// queryPaging đọc page & page_size (mặc định 1/20).
func queryPaging(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	return page, pageSize
}
