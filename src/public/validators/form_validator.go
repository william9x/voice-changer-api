package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

func RegisterFormValidators() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("notblank", validators.NotBlank); err != nil {
			return err
		}
		if err := v.RegisterValidation("tasktype", tasktype); err != nil {
			return err
		}
	}
	return nil
}

func tasktype(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return val == "vc:rvc" || val == "aic"
}
