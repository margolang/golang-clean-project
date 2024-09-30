package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"presentation/internal/api"
	"presentation/internal/domain/application/errors"
)

func ServError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	switch err.(type) {
	case validator.ValidationErrors:
		api.RespondValidationErrors(c, err)
	default:
		switch err {
		case errors.ErrApplicationNotFound:
			api.RespondNotFound(c, err)
		default:
			log.Println("application internal error: %v", err)
			api.RespondInternalError(c, api.ErrInternalServerError.Error(), err)
		}
	}
}
