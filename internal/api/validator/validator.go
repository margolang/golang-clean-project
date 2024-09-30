package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"regexp"
	"strings"
)

var (
	inputNameValidator, _ = regexp.Compile("^[a-zA-Z\\s-]+$")
)

var inputNameFormat validator.Func = func(fl validator.FieldLevel) bool {
	inputText, ok := fl.Field().Interface().(string)
	if ok {
		if inputText == "Shumaher" {
			return false
		}
	}

	return true
}

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

		err := v.RegisterValidation("inputNameFormat", inputNameFormat)
		if err != nil {
			log.Fatal("Failed to register inputNameFormat validator")
		}
	}
}
