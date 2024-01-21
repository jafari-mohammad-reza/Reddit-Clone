package setuproutes

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/api/handler"
	"github.com/reddit-clone/src/api/middlewares"
	"github.com/reddit-clone/src/share/config"
)

func registerRoutes(r *gin.Engine, cfg *config.Config) {

	// api := r.Group("/api")
	// apiGroup = v1
	// v1.Use(middlewares.LoggerMiddleware(logger))
	limiter := middlewares.NewLimmiterMiddlware(cfg)
	r.Use(limiter.RateLimiter())
	r.Use(middlewares.ResponseFormatterMiddleware())
	r.POST("/Debug/test", handler.CreateHandler)
}
