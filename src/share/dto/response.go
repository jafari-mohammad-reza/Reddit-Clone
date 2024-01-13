package dto

import "time"

type BaseResponse struct {
	StatusCode int       `json:"statusCode"`
	TimeSpan   time.Time `json:"timeSpan"`
}

type SuccessResponse[T any] struct {
	BaseResponse
	Result   T           `json:"result"`
	Metadata interface{} `json:"metadata"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	Path         string `json:"path"`
	BaseResponse
}

