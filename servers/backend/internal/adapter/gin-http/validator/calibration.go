package validator

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func calibrationValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation(
			"five_point",
			func(fl validator.FieldLevel) bool {
				regex, compileErr := regexp.Compile(`^([-]?\d*\.?\d+\,){4}[-]?\d*\.?\d+$`)
				if compileErr != nil {
					return false
				}
				return regex.MatchString(fl.Field().String())
			})
		if err != nil {
			log.Panic(err)
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation(
			"five_point_json",
			func(fl validator.FieldLevel) bool {
				regex, compileErr := regexp.Compile(`^{"value": "([-]?\d*\.?\d+\,){4}[-]?\d*\.?\d+", "parameter": "([-]?\d*\.?\d+\,){4}[-]?\d*\.?\d+"}$`)
				if compileErr != nil {
					return false
				}
				return regex.MatchString(fl.Field().String())
			})
		if err != nil {
			log.Panic(err)
		}
	}
}
