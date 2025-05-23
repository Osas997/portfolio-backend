package errorhandler

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, e := range err {
		field := strings.ToLower(e.Field())
		errors[field] = msgForTag(e.Tag())
	}
	return errors
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "field is required"
	case "email":
		return "invalid email format"
	case "min":
		return "value is too short"
	case "max":
		return "value is too long"
	case "isFile":
		return "field must be a file"
	case "image":
		return "field must be an image file"
	case "fileSize":
		return "file size is too large"
	default:
		return "invalid value"
	}
}
