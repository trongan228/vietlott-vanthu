package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

// TimKiemVanBanHandler tìm kiếm văn bản theo tên (trích yếu) và số văn bản
// trên cả 3 bảng cùng lúc.
type TimKiemVanBanHandler struct {
	svc *service.TimKiemVanBanService
}

func NewTimKiemVanBanHandler(svc *service.TimKiemVanBanService) *TimKiemVanBanHandler {
	return &TimKiemVanBanHandler{svc: svc}
}

func (h *TimKiemVanBanHandler) Register(rg *gin.RouterGroup) {
	rg.GET("/van-ban/tim-kiem", h.TimKiem)
}

// TimKiem: GET /van-ban/tim-kiem?q=báo cáo&so=123&nam=2026&limit=20
// Cần ít nhất một trong q hoặc so.
func (h *TimKiemVanBanHandler) TimKiem(c *gin.Context) {
	var p model.TimKiemVanBanParams

	p.Q = c.Query("q")
	p.So = c.Query("so")

	nam, ok := queryInt16(c, "nam")
	if !ok {
		return
	}
	p.Nam = nam
	p.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "20"))

	result, err := h.svc.TimKiem(c.Request.Context(), p)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
