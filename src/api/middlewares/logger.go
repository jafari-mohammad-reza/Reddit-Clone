package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/reddit-clone/src/api/helper"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)

func LoggerMiddleware(lg custome_logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		req := ctx.Request
		wrapper := &api.ResponseWriterWrapper{ResponseWriter: ctx.Writer}
		ctx.Writer = wrapper
		path := req.URL.Path
		query := req.URL.RawQuery
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
		method := req.Method
		ctx.Next()

		status := wrapper.StatusCode
		body := wrapper.Body.String()
		end := time.Since(start)
		log := make(map[custome_logger.ExtraKey]interface{})
		log["Path"] = path
		log["Query"] = query
		log["Method"] = method
		log["Status"] = status
		log["Body"] = body
		log["Duration"] = end
		lg.Info(custome_logger.API, custome_logger.RequestLog, "", log)
	}
}
