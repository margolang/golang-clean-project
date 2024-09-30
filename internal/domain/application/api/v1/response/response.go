package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"presentation/internal/api"
	domain "presentation/internal/domain/application"
)

type ApplicationResponse struct {
	Status     string      `json:"status"`
	HTTPStatus int         `json:"http_status"`
	Data       application `json:"data"`
} // @name Application

func RespondApplication(c *gin.Context, locale string, application domain.Application) {
	api.RespondJSON(
		c, http.StatusOK, locale, ApplicationResponse{
			Status:     "success",
			HTTPStatus: http.StatusOK,
			Data:       ApplicationConverter.Response(application),
		},
	)
}
