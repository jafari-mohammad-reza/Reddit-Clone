package subreddit

import (
	"Reddit-Clone/src/share/config"
	"Reddit-Clone/src/share/database/cache"
	"Reddit-Clone/src/share/pkg/custome_logger"
	"Reddit-Clone/src/share/services"
)

type SubredditService struct {
	cfg             *config.Config
	redisClient     *redis.Client
	rabbitMqService *services.RabbitMQService
}

func NewSubredditService(cfg *config.Config) *SubredditService {
	lg := custome_logger.NewLogger(cfg)
	return &SubredditService{
		cfg:             cfg,
		redisClient:     cache.GetRedisClient(),
		rabbitMqService: services.NewRabbitMQService(cfg, lg, "category", nil), // this setting just for now
	}
}
