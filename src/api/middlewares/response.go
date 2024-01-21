package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/reddit-clone/src/api/helper"
	"github.com/reddit-clone/src/share/dto"
)

func ResponseFormatterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		w := &api.ResponseWriterWrapper{ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		response := dto.SuccessResponse[any]{
			BaseResponse: dto.BaseResponse{
				StatusCode: w.StatusCode,
				TimeSpan:   time.Now(),
			},
			Result:   json.RawMessage(w.Body.Bytes()),
			Metadata: nil,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(w.StatusCode)
		err := json.NewEncoder(c.Writer).Encode(response)
		if err != nil {
			errorResponse := dto.ErrorResponse{
				ErrorMessage: err.Error(),
				Path:         c.Request.RequestURI,
				BaseResponse: dto.BaseResponse{},
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
	}
}
