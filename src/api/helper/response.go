package api

import (
	"net/http"
	"time"

	"github.com/reddit-clone/src/share/dto"
)

func GenerateSuccessResponse[T any](response T, status *int, metaData *interface{}) *dto.SuccessResponse[T] {
	var statusCode int
	if status != nil {
		statusCode = *status
	} else {
		statusCode = http.StatusOK
	}
	return &dto.SuccessResponse[T]{
		BaseResponse: dto.BaseResponse{
			StatusCode: statusCode,
			TimeSpan:   time.Now(),
		},
		Result:   response,
		Metadata: metaData,
	}
}

func GenerateErrorResponse(err error , path string,status *int) *dto.ErrorResponse {
	var statusCode int
	if status != nil {
		statusCode = *status
	} else {
		statusCode = http.StatusOK
	}
	return &dto.ErrorResponse{
		ErrorMessage: err.Error(),
		Path: path,
		BaseResponse: dto.BaseResponse{
			StatusCode: statusCode,
			TimeSpan:   time.Now(),
		},
	}
}