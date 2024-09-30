package i18n

import (
	//"dev.azure.com/competommc/Marketplace/mp-common.git/api"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	api "presentation/internal/api/v1"
	"sync"
)

//go:embed locales/locale.*.json
var localeFS embed.FS

var (
	defaultErrorKey = "error.default_message"
	once            sync.Once
	localeMap       map[Locale]localeDictionary
)

type Locale string

type localeDictionary map[string]string

type Config struct {
	Locales       []Locale
	UnmarshalFunc func([]byte, interface{}) error
}

func newLocales(cfg Config) {
	if localeMap == nil {
		localeMap = make(map[Locale]localeDictionary, len(cfg.Locales))
	}

	for _, lc := range cfg.Locales {
		data, err := localeFS.ReadFile(fmt.Sprintf("locales/locale.%s.json", lc))
		if err != nil {
			panic(err)
		}

		messages := make(map[string]string, 100)
		if err := cfg.UnmarshalFunc(data, &messages); err != nil {
			panic(err)
		}

		localeMap[lc] = messages
	}
}

func LocaleMiddleware(cfg Config) gin.HandlerFunc {
	if len(cfg.Locales) == 0 {
		cfg.Locales = []Locale{api.DefaultLocale}
	}

	if cfg.UnmarshalFunc == nil {
		cfg.UnmarshalFunc = json.Unmarshal
	}

	once.Do(func() {
		newLocales(cfg)
	})

	return func(c *gin.Context) {
		api.SetLocale(c)
		c.Next()
	}
}

func Localize(c *gin.Context, key string, param string) string {
	locale := (Locale)(api.Locale(c))

	if localeMap == nil {
		return key
	}

	localeDic, ok := localeMap[locale]
	if !ok {
		return key
	}

	value, ok := localeDic[key]
	if !ok {
		return key
	}
	if param != "" {
		value = fmt.Sprintf(value, param)
	}

	return value
}

func LocalizeError(c *gin.Context, err error) error {
	var errMsg string
	if err != nil {
		errMsg = Localize(c, err.Error(), "")
	} else {
		errMsg = Localize(c, defaultErrorKey, "")
	}

	if errMsg == "" || errMsg == err.Error() {
		return err
	}

	return errors.New(errMsg)
}
