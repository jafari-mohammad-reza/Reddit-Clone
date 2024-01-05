package category

import (
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
	"github.com/reddit-clone/src/share/services"
	"github.com/redis/go-redis/v9"
)

type CategoryService struct {
	cfg             *config.Config
	redisClient     *redis.Client
	rabbitMqService *services.RabbitMQService
}

func NewCategoryService(cfg *config.Config) *CategoryService {
	lg := custome_logger.NewLogger(cfg)
	return &CategoryService{
		cfg:             cfg,
		redisClient:     cache.GetRedisClient(),
		rabbitMqService: services.NewRabbitMQService(cfg, lg, "category", nil), // this setting just for now
	}
}
