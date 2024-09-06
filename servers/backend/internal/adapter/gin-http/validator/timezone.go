package validator

import (
	"log"

	"auto-monitoring/pkg/regexp_check"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func timezoneValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation(
			"timezone",
			func(fl validator.FieldLevel) bool {
				return regexp_check.IsTimeZone(fl.Field().String())
			})
		if err != nil {
			log.Panic(err)
		}
	}
}
