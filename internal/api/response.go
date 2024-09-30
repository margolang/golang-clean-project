package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"presentation/i18n"
	api "presentation/internal/api/v1"
	domainvalidator "presentation/internal/api/validator"
)

const (
	HeaderContentTypeOptions = "X-Content-Type-Options"
	HeaderAuthenticate       = "WWW-Authenticate"
	HeaderContentLanguage    = "Content-Language"
)

type ErrorContextMessage struct {
	Message string `json:"message"`
} // @name ErrorContextMessage

type ErrorContext map[string][]ErrorContextMessage // @name ErrorContext

func NewErrorContext(c *gin.Context, validationErrors validator.ValidationErrors) map[string]ErrorContext {
	if validationErrors == nil {
		return nil
	}

	errorContext := make(ErrorContext, len(validationErrors))
	for _, err := range validationErrors {
		errorContext[err.Field()] = []ErrorContextMessage{{
			Message: i18n.Localize(c, err.Error(), err.Param()),
		}}
	}

	return map[string]ErrorContext{"errors": errorContext}
}

type Response struct {
} // @name Response

func RespondCreated(c *gin.Context) {
	RespondJSON(c, http.StatusCreated, api.Locale(c), Response{})
}

func RespondNoContent(c *gin.Context) {
	RespondJSON(c, http.StatusNoContent, api.Locale(c), Response{})
}

func RespondAccepted(c *gin.Context) {
	RespondJSON(c, http.StatusAccepted, api.Locale(c), Response{})
}

func RespondOK(c *gin.Context, data any) {
	RespondJSON(c, http.StatusOK, api.Locale(c), data)
}

type ErrorResponse struct {
	Type    string                  `json:"type"`
	Message string                  `json:"message"`
	Context map[string]ErrorContext `json:"context,omitempty"`
} // @name ErrorResponse

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", r.Type, r.Message)
}

func NewErrorResponse(c *gin.Context, err error) ErrorResponse {
	var bindingError validator.ValidationErrors
	if errors.As(err, &bindingError) {
		return ErrorResponse{
			Type:    ErrValidation.Error(),
			Message: descriptionError(c, ErrValidation),
			Context: NewErrorContext(c, bindingError),
		}
	}

	return ErrorResponse{
		Type:    err.Error(),
		Message: descriptionError(c, err),
	}
}

func RespondBadRequest(c *gin.Context, err error) {
	respondErrorJSON(c, http.StatusBadRequest, err, NewErrorResponse(c, err))
}

func RespondValidationErrors(c *gin.Context, err error) {
	validationErr := domainvalidator.ToValidationContext(err)
	respondErrorJSON(c, http.StatusBadRequest, err, NewErrorResponse(c, validationErr))
}

func RespondTooManyRequest(c *gin.Context, err error) {
	respondErrorJSON(c, http.StatusTooManyRequests, err, NewErrorResponse(c, err))
}

func RespondUnauthorized(c *gin.Context, err error) {
	c.Header(HeaderAuthenticate, "Bearer realm=Access to API")
	respondErrorJSON(c, http.StatusUnauthorized, err, NewErrorResponse(c, err))
}

func RespondForbidden(c *gin.Context, err error) {
	respondErrorJSON(c, http.StatusForbidden, err, NewErrorResponse(c, err))
}

func RespondNotFound(c *gin.Context, err error) {
	respondErrorJSON(c, http.StatusNotFound, err, NewErrorResponse(c, err))
}

func RespondConflict(c *gin.Context, err error) {
	respondErrorJSON(c, http.StatusConflict, err, NewErrorResponse(c, err))
}

func RespondInternalError(c *gin.Context, errorType string, err error) {
	respondErrorJSON(c, http.StatusInternalServerError, err, ErrorResponse{
		Type:    errorType,
		Message: descriptionError(c, err),
	})
}

func respondErrorJSON(c *gin.Context, s int, err error, d interface{}) {
	if err != nil {
		c.Error(err) // nolint: errcheck
	}
	c.Header(HeaderContentTypeOptions, "nosniff")
	c.AbortWithStatusJSON(s, d)
}

func descriptionError(c *gin.Context, err error) string {
	return i18n.LocalizeError(c, err).Error()
}

func RespondJSON(c *gin.Context, s int, l string, d interface{}) {
	c.Header(HeaderContentTypeOptions, "nosniff")
	c.Header(HeaderContentLanguage, l)
	c.JSON(s, d)
}
