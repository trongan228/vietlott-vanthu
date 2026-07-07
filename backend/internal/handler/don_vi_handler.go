package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

type DonViHandler struct {
	svc *service.DonViService
}

func NewDonViHandler(svc *service.DonViService) *DonViHandler {
	return &DonViHandler{svc: svc}
}

func (h *DonViHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/don-vi")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *DonViHandler) Create(c *gin.Context) {
	var in model.DonViInput
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

func (h *DonViHandler) GetByID(c *gin.Context) {
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

func (h *DonViHandler) Update(c *gin.Context) {
	id, ok := parseIDParam32(c)
	if !ok {
		return
	}

	var in model.DonViInput
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

func (h *DonViHandler) Delete(c *gin.Context) {
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

func (h *DonViHandler) List(c *gin.Context) {
	var f model.DonViFilter

	if v := c.Query("loai_don_vi"); v != "" {
		f.LoaiDonVi = &v
	}
	isActive, ok := queryBool(c, "is_active")
	if !ok {
		return
	}
	f.IsActive = isActive
	f.Search = c.Query("q")
	f.Page, f.PageSize = queryPaging(c)

	result, err := h.svc.List(c.Request.Context(), f)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
