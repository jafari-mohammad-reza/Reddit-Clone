package category

import (
	"Reddit-Clone/src/share/config"
	"Reddit-Clone/src/share/database/cache"
	"Reddit-Clone/src/share/pkg/custome_logger"
	"Reddit-Clone/src/share/services"
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
