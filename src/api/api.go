package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/docs"
	"github.com/reddit-clone/src/domains/routes"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
	"github.com/reddit-clone/src/share/services"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = custome_logger.NewLogger(config.GetConfig())
var apiGroup *gin.RouterGroup

type CategoryService struct {
	cfg             *config.Config
	redisClient     *redis.Client
	rabbitMqService *services.RabbitMQService
	pgRepository    *sql.DB
}

type createCategoryDto struct {
	Name         string `json:"name,omitemty" binding:"required" , oneof="Hot New Top Rising"`
	CategoryType string `json:"category_type,omitempty" binding:"required" oneof="Hot New Top Rising"`
}
type Category struct {
	_id          string
	categoryType string
}

func NewCategoryService(cfg *config.Config) *CategoryService {
	lg := custome_logger.NewLogger(cfg)
	pg := postgres.GetPostgres()
	return &CategoryService{
		cfg:             cfg,
		redisClient:     cache.GetRedisClient(),
		pgRepository:    pg,
		rabbitMqService: services.NewRabbitMQService(cfg, lg, "category", nil), // this setting just for now
	}
}
func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routes.RegisterRoutes(r, cfg)
	registerSwagger(r, cfg)
	logger.Info(custome_logger.General, custome_logger.Startup, "Server started", nil)
	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func GetApiRoute() *gin.RouterGroup {
	return apiGroup
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

// Category SUBreddit => hot / new / top / rising
// search system for finding increasing post (Popularity!!! karma point)
// write sql query
