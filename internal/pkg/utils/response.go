package utils

type WebResponse struct {
	Message string          `json:"message"`
	Data    any             `json:"data,omitempty"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}

func NewWebResponse(message string, data any) *WebResponse {
	return &WebResponse{Message: message, Data: data}
}

type PaginationMeta struct {
	Page            int   `json:"page"`
	Limit           int   `json:"limit"`
	ProductsCount   int64 `json:"productsCount"`
	PageCount       int   `json:"pageCount"`
	HasPreviousPage bool  `json:"hasPreviousPage"`
	HasNextPage     bool  `json:"hasNextPage"`
}

func NewPaginationMeta(page, limit int, productsCount int64) *PaginationMeta {
	pageCount := 0
	if limit > 0 {
		pageCount = int((productsCount + int64(limit) - 1) / int64(limit))
	}
	hasPrev := page > 1
	hasNext := page < pageCount
	return &PaginationMeta{
		Page:            page,
		Limit:           limit,
		ProductsCount:   productsCount,
		PageCount:       pageCount,
		HasPreviousPage: hasPrev,
		HasNextPage:     hasNext,
	}
}
