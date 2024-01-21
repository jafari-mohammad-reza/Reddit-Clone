package category

import (
	"database/sql"

	"github.com/reddit-clone/src/domains/subreddit-domain/category/dtos"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
	"github.com/reddit-clone/src/share/services"
	"github.com/redis/go-redis/v9"
)

type CategoryService struct {
	cfg             *config.Config
	redisClient     *redis.Client
	rabbitMqService *services.RabbitMQService
	pgRepository    *sql.DB
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

// Category SUBreddit => hot / new / top / rising
// search system for finding increasing post (Popularity!!! karma point)
// write sql query
type Category struct {
	_id          string
	category     *Category
	categoryType string `json:"category_type,omitempty" binding:"required,currency" oneof="Hot New Top Rising"`
}

func (s *CategoryService) Create(dto dtos.CreateCategoryDto) (sql.Result, error) {
	category, err := s.pgRepository.Exec("INSERT INTO category (dto.categoryType, name,ParentCategory ,) VALUES ($1, $2 )", dto.CategoryType, dto.Name)
	if err != nil {
		return nil, err
	}

	return category, nil
}
