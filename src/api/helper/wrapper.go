package api

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body       bytes.Buffer
	StatusCode int
}

func (w *ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
