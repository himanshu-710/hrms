package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var errs []string
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
	}
	return fmt.Errorf("%s", strings.Join(errs, ", "))
}
