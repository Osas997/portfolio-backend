package utils

type WebResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewWebResponse(message string, data any) *WebResponse {
	return &WebResponse{Message: message, Data: data}
}
