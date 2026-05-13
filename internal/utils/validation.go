package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) map[string]string{
	if errs, ok := err.(validator.ValidationErrors); ok{
		res := make(map[string]string)

		for _, err := range errs{
			switch err.Tag(){
			case "required":
				res[err.Field()] = fmt.Sprintf("%s is required", err.Field())
			case "email":
				res[err.Field()] = fmt.Sprintf("%s must be a valid email", err.Field())
			case "min":
				res[err.Field()] = fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
			case "max":
				res[err.Field()] = fmt.Sprintf("%s must be at most %s characters", err.Field(), err.Param())
			default:
				res[err.Field()] = fmt.Sprintf("%s is invalid", err.Field())
			}
		}
		return res
	}

	return nil

}