package dto

type PaginationRequest struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit"`
}

type PaginationResponse[T any] struct {
	Response    SuccessResponse[T] `json:"response"`
	Page        int                `json:"page,omitempty"`
	Limit       int                `json:"limit"`
	PageNumbers int                `json:"pageNumbers"`
	ItemCount   int                `json:"itemCount"`
}
