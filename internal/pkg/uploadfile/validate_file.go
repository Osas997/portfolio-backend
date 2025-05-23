package uploadfile

import (
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterValidation("isFile", validateFile)
	v.RegisterValidation("image", validateImage)
	v.RegisterValidation("fileSize", validateFileSize)
}

func validateFile(fl validator.FieldLevel) bool {
	field := fl.Field()

	if fileHeader, ok := field.Interface().(multipart.FileHeader); ok {
		return fileHeader.Size > 0
	}

	return true
}

func validateImage(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	// Allowed extensions
	allowedExts := []string{".jpg", ".jpeg", ".png"}
	filename := strings.ToLower(file.Filename)
	for _, ext := range allowedExts {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func validateFileSize(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param() // Ini akan berisi angka dalam MB
	maxSizeMB, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return false
	}

	maxSize := int64(maxSizeMB * 1024 * 1024)

	if fileHeader, ok := field.Interface().(multipart.FileHeader); ok {
		return fileHeader.Size <= maxSize
	}

	return true
}
