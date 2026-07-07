package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

type HoSoLuuTruHandler struct {
	svc *service.HoSoLuuTruService
}

func NewHoSoLuuTruHandler(svc *service.HoSoLuuTruService) *HoSoLuuTruHandler {
	return &HoSoLuuTruHandler{svc: svc}
}

func (h *HoSoLuuTruHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/ho-so-luu-tru")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *HoSoLuuTruHandler) Create(c *gin.Context) {
	var in model.HoSoLuuTruInput
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

func (h *HoSoLuuTruHandler) GetByID(c *gin.Context) {
	id, ok := parseIDParam32(c)
	if !ok {
		return
	}

	m, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *HoSoLuuTruHandler) Update(c *gin.Context) {
	id, ok := parseIDParam32(c)
	if !ok {
		return
	}

	var in model.HoSoLuuTruInput
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

func (h *HoSoLuuTruHandler) Delete(c *gin.Context) {
	id, ok := parseIDParam32(c)
	if !ok {
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *HoSoLuuTruHandler) List(c *gin.Context) {
	var f model.HoSoLuuTruFilter

	hopID, ok := queryInt32(c, "hop_id")
	if !ok {
		return
	}
	f.HopID = hopID

	if v := c.Query("loai_tap"); v != "" {
		f.LoaiTap = &v
	}

	nam, ok := queryInt16(c, "nam")
	if !ok {
		return
	}
	f.Nam = nam

	vinhVien, ok := queryBool(c, "vinh_vien")
	if !ok {
		return
	}
	f.VinhVien = vinhVien

	f.Search = c.Query("q")
	f.Page, f.PageSize = queryPaging(c)

	result, err := h.svc.List(c.Request.Context(), f)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
