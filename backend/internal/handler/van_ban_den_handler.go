package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

type VanBanDenHandler struct {
	svc *service.VanBanDenService
}

func NewVanBanDenHandler(svc *service.VanBanDenService) *VanBanDenHandler {
	return &VanBanDenHandler{svc: svc}
}

func (h *VanBanDenHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/van-ban-den")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *VanBanDenHandler) Create(c *gin.Context) {
	var in model.VanBanDenInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := h.svc.Create(c.Request.Context(), &in)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusCreated, m)
}

func (h *VanBanDenHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	m, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *VanBanDenHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	var in model.VanBanDenInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := h.svc.Update(c.Request.Context(), id, &in)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *VanBanDenHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *VanBanDenHandler) List(c *gin.Context) {
	var f model.VanBanDenFilter

	if v := c.Query("nam"); v != "" {
		n, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tham số nam không hợp lệ"})
			return
		}
		nn := int16(n)
		f.Nam = &nn
	}
	if v := c.Query("so_den"); v != "" {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tham số so_den không hợp lệ"})
			return
		}
		nn := int32(n)
		f.SoDen = &nn
	}
	if v := c.Query("so_ky_hieu"); v != "" {
		f.SoKyHieu = &v
	}

	noiGuiID, ok := queryInt32(c, "noi_gui_id")
	if !ok {
		return
	}
	f.NoiGuiID = noiGuiID

	donViXuLyID, ok := queryInt32(c, "don_vi_xu_ly_id")
	if !ok {
		return
	}
	f.DonViXuLyID = donViXuLyID

	f.Search = c.Query("q")
	f.SortBy = c.Query("sort_by")
	f.SortDir = c.Query("sort_dir")
	f.Page, f.PageSize = queryPaging(c)

	result, err := h.svc.List(c.Request.Context(), f)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
