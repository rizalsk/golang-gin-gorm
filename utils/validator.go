package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetValidation(err error) string {
	var errStr string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		field := strings.ToLower(string(fieldErr.StructField()))
		// tag := fieldErr.ActualTag()
		// errStr := fmt.Sprintf("The %s error in %s", field, tag)
		errSplit := strings.Split(fieldErr.Error(), ":")
		errStr = strings.Replace(errSplit[2], fieldErr.StructField(), field, -1)
	}

	return errStr
}
