package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vanthu-backend/internal/httpx"
	"vanthu-backend/internal/model"
	"vanthu-backend/internal/service"
)

type HopHandler struct {
	svc *service.HopService
}

func NewHopHandler(svc *service.HopService) *HopHandler {
	return &HopHandler{svc: svc}
}

func (h *HopHandler) Register(rg *gin.RouterGroup) {
	g := rg.Group("/hop")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *HopHandler) Create(c *gin.Context) {
	var in model.HopInput
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

func (h *HopHandler) GetByID(c *gin.Context) {
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

func (h *HopHandler) Update(c *gin.Context) {
	id, ok := parseIDParam32(c)
	if !ok {
		return
	}

	var in model.HopInput
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

func (h *HopHandler) Delete(c *gin.Context) {
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

func (h *HopHandler) List(c *gin.Context) {
	var f model.HopFilter

	soHop, ok := queryInt32(c, "so_hop")
	if !ok {
		return
	}
	f.SoHop = soHop

	if v := c.Query("loai_hop"); v != "" {
		f.LoaiHop = &v
	}

	thungID, ok := queryInt32(c, "thung_id")
	if !ok {
		return
	}
	f.ThungID = thungID
	f.Page, f.PageSize = queryPaging(c)

	result, err := h.svc.List(c.Request.Context(), f)
	if err != nil {
		httpx.WriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
