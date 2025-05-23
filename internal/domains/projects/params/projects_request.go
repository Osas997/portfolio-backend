package params

import "mime/multipart"

type ProjectRequest struct {
	Title     string                  `form:"title" validate:"required,min=3"`
	Content   string                  `form:"content" validate:"required,min=3"`
	Thumbnail *multipart.FileHeader   `form:"thumbnail" validate:"required,isFile,image,fileSize=2"`
	Images    []*multipart.FileHeader `form:"images" validate:"omitempty,dive,isFile,image,fileSize=2"`
}

type UpdateProjectRequest struct {
	Title     *string                 `form:"title" validate:"omitempty,min=3"`
	Content   *string                 `form:"content" validate:"omitempty,min=3"`
	Thumbnail *multipart.FileHeader   `form:"thumbnail" validate:"omitempty,isFile,image,fileSize=2"`
	Images    []*multipart.FileHeader `form:"images" validate:"omitempty,dive,isFile,image,fileSize=2"`
}
