package uploadfile

import (
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

func FormatUpload(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)

	uniqueFileName := uuid.New().String() + ext

	return uniqueFileName
}
