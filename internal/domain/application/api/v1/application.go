package v1

import (
	"github.com/gin-gonic/gin"
	baseapi "presentation/internal/api"
	basev1 "presentation/internal/api/v1"
	domain "presentation/internal/domain/application"
	"presentation/internal/domain/application/api"
	"presentation/internal/domain/application/api/v1/response"
)

type applicationHandler struct {
	service domain.ApplicationService
}

func newApplicationHandler(s domain.ApplicationService) *applicationHandler {
	return &applicationHandler{service: s}
}

func (h *applicationHandler) Create(c *gin.Context) {
	var req applicationReq
	if err := c.ShouldBind(&req); err != nil {
		api.ServError(c, err)
		return
	}

	err := h.service.Create(c, RequestConverter.ApplicationCreate(req))
	if err != nil {
		api.ServError(c, err)
		return
	}

	baseapi.RespondNoContent(c)
}

func (h *applicationHandler) GetByID(c *gin.Context) {
	applicationID := c.Param(basev1.ApplicationID)

	res, err := h.service.Get(c, applicationID)
	if err != nil {
		api.ServError(c, err)
		return
	}

	response.RespondApplication(c, "ru", res)
}

func (h *applicationHandler) Delete(c *gin.Context) {
	applicationID := c.Param(basev1.ApplicationID)

	err := h.service.Delete(c, applicationID)
	if err != nil {
		api.ServError(c, err)
		return
	}

	baseapi.RespondNoContent(c)
}

func RegisterApplicationHandler(r *gin.RouterGroup, s domain.ApplicationService) {
	h := newApplicationHandler(s)

	g := r.Group("/applications")
	{
		g.POST("", h.Create)
		g.GET("/:application_id", h.GetByID)
		g.DELETE("/:application_id", h.Create)
	}
}
