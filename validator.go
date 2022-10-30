package validatorformater

import (
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FormatErrorValidation(err error, message string) ErrorResponse {
	var errorResponse ErrorResponse

	var errors []string
	tempError := ""
	for _, v := range err.(validator.ValidationErrors) {
		splitted := camelcase.Split(v.Field())
		field := strings.Join(splitted, "_")
		switch v.Tag() {
		case "required":
			tempError = fmt.Sprintf("%s is required", field)
		case "email":
			tempError = fmt.Sprintf("%s is not valid email", field)
		case "gte":
			tempError = fmt.Sprintf("%s value must be greater than %s", field, v.Param())
		case "lte":
			tempError = fmt.Sprintf("%s value must be lower than %s", field, v.Param())
		case "min":
			tempError = fmt.Sprintf("%s character must be min %s", field, v.Param())
		case "max":
			tempError = fmt.Sprintf("%s character must be max %s", field, v.Param())
		default:
			tempError = v.Error()

		}
		errors = append(errors, tempError)
	}
	errorResponse.Status = "Bad Request"
	errorResponse.Message = message
	errorResponse.Data = errors

	return errorResponse
}
