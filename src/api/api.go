package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/docs"
	"github.com/reddit-clone/src/api/middlewares"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = custome_logger.NewLogger(config.GetConfig())
var apiGroup *gin.RouterGroup

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	registerRoutes(r, cfg)
	registerSwagger(r, cfg)
	logger.Info(custome_logger.General, custome_logger.Startup, "Server started", nil)
	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func GetApiRoute() *gin.RouterGroup {
	return apiGroup
}
func registerRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	apiGroup = v1
	v1.Use(middlewares.LoggerMiddleware(logger))
	limiter := middlewares.NewLimmiterMiddlware(cfg)
	v1.Use(limiter.RateLimiter())
	v1.Use(middlewares.ResponseFormatterMiddleware())
}

func registerSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Reddit Clone"
	docs.SwaggerInfo.Description = "reddit clone api made using gin & swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
