package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

type CanBoHandler struct {
	svc *service.CanBoService
}

func NewCanBoHandler(svc *service.CanBoService) *CanBoHandler {
	return &CanBoHandler{svc: svc}
}

func (h *CanBoHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/can-bo")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *CanBoHandler) Create(c *gin.Context) {
	var in model.CanBoInput
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

func (h *CanBoHandler) GetByID(c *gin.Context) {
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

func (h *CanBoHandler) Update(c *gin.Context) {
	id, ok := parseIDParam32(c)
	if !ok {
		return
	}

	var in model.CanBoInput
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

func (h *CanBoHandler) Delete(c *gin.Context) {
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

func (h *CanBoHandler) List(c *gin.Context) {
	var f model.CanBoFilter

	donViID, ok := queryInt32(c, "don_vi_id")
	if !ok {
		return
	}
	f.DonViID = donViID

	isVanThu, ok := queryBool(c, "is_van_thu")
	if !ok {
		return
	}
	f.IsVanThu = isVanThu

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
