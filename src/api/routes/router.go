package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/api/middlewares"
	category "github.com/reddit-clone/src/domains/subreddit-domain/categoryRoutes"
	"github.com/reddit-clone/src/share/config"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	// api := r.Group("/api")
	// apiGroup = v1
	// v1.Use(middlewares.LoggerMiddleware(logger))
	limiter := middlewares.NewLimmiterMiddlware(cfg)
	r.Use(limiter.RateLimiter())
	r.Use(middlewares.ResponseFormatterMiddleware())
	setupRoutes(r)

}

func setupRoutes(server *gin.Engine) {
	api := server.Group("api")

	category.CategoryRoutes(api)
}
