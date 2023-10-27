package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
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

func ParseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {

		case validator.ValidationErrors:
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}
	return out
}

func parseFieldError(e validator.FieldError) string {
	fieldPrefix := strings.ToLower(fmt.Sprintf("The field %s", e.Field()))
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required":
		return fmt.Sprintf("%s is %s", fieldPrefix, tag)
	case "numeric":
		return fmt.Sprintf("%s is not %s", fieldPrefix, tag)
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	default:
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		} else {
			return fmt.Errorf("%v", e).Error()
		}
	}
}
func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}
