package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

type VanBanDiKhongTenLoaiHandler struct {
	svc *service.VanBanDiKhongTenLoaiService
}

func NewVanBanDiKhongTenLoaiHandler(svc *service.VanBanDiKhongTenLoaiService) *VanBanDiKhongTenLoaiHandler {
	return &VanBanDiKhongTenLoaiHandler{svc: svc}
}

func (h *VanBanDiKhongTenLoaiHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/van-ban-di-khong-ten-loai")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *VanBanDiKhongTenLoaiHandler) Create(c *gin.Context) {
	var in model.VanBanDiKhongTenLoaiInput
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

func (h *VanBanDiKhongTenLoaiHandler) GetByID(c *gin.Context) {
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

func (h *VanBanDiKhongTenLoaiHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id không hợp lệ"})
		return
	}

	var in model.VanBanDiKhongTenLoaiInput
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

func (h *VanBanDiKhongTenLoaiHandler) Delete(c *gin.Context) {
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

func (h *VanBanDiKhongTenLoaiHandler) List(c *gin.Context) {
	var f model.VanBanDiKhongTenLoaiFilter

	if v := c.Query("nam"); v != "" {
		n, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tham số nam không hợp lệ"})
			return
		}
		nn := int16(n)
		f.Nam = &nn
	}
	if v := c.Query("so_ky_hieu"); v != "" {
		f.SoKyHieu = &v
	}

	nguoiKyID, ok := queryInt32(c, "nguoi_ky_id")
	if !ok {
		return
	}
	f.NguoiKyID = nguoiKyID

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
