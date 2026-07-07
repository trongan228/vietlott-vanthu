package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/service"
)

// ViTriVanBanHandler tra cứu vị trí lưu kho của văn bản
// (văn bản → hồ sơ → hộp → thùng) từ 3 view v_vi_tri_van_ban_*.
type ViTriVanBanHandler struct {
	svc *service.ViTriVanBanService
}

func NewViTriVanBanHandler(svc *service.ViTriVanBanService) *ViTriVanBanHandler {
	return &ViTriVanBanHandler{svc: svc}
}

func (h *ViTriVanBanHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/vi-tri")
	g.GET("/van-ban-den", h.VanBanDen)
	g.GET("/van-ban-di-co-ten-loai", h.VanBanDiCoTenLoai)
	g.GET("/van-ban-di-khong-ten-loai", h.VanBanDiKhongTenLoai)
}

// VanBanDen: GET /vi-tri/van-ban-den?nam=2022&so_den=150
func (h *ViTriVanBanHandler) VanBanDen(c *gin.Context) {
	nam, ok := queryInt16(c, "nam")
	if !ok {
		return
	}
	soDen, ok := queryInt32(c, "so_den")
	if !ok {
		return
	}

	items, err := h.svc.TraCuuVanBanDen(c.Request.Context(), nam, soDen)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// VanBanDiCoTenLoai: GET /vi-tri/van-ban-di-co-ten-loai?nam=2022&so_ky_hieu=...
func (h *ViTriVanBanHandler) VanBanDiCoTenLoai(c *gin.Context) {
	nam, ok := queryInt16(c, "nam")
	if !ok {
		return
	}
	var soKyHieu *string
	if v := c.Query("so_ky_hieu"); v != "" {
		soKyHieu = &v
	}

	items, err := h.svc.TraCuuVanBanDiCoTenLoai(c.Request.Context(), nam, soKyHieu)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// VanBanDiKhongTenLoai: GET /vi-tri/van-ban-di-khong-ten-loai?nam=2022&so_ky_hieu=...
func (h *ViTriVanBanHandler) VanBanDiKhongTenLoai(c *gin.Context) {
	nam, ok := queryInt16(c, "nam")
	if !ok {
		return
	}
	var soKyHieu *string
	if v := c.Query("so_ky_hieu"); v != "" {
		soKyHieu = &v
	}

	items, err := h.svc.TraCuuVanBanDiKhongTenLoai(c.Request.Context(), nam, soKyHieu)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
