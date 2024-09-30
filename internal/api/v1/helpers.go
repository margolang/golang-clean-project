package v1

import "github.com/gin-gonic/gin"

const (
	HeaderAcceptLanguage = "Accept-Language"
	DefaultLocale        = "en"
	KeyLocale            = "locale"
)

func SetLocale(c *gin.Context) {
	locale := c.GetHeader(HeaderAcceptLanguage)
	if locale == "" {
		locale = DefaultLocale
	}

	c.Set(KeyLocale, locale)
}

func Locale(c *gin.Context) string {
	l, ok := c.Get(KeyLocale)
	if !ok {
		return DefaultLocale
	}

	return l.(string)
}
