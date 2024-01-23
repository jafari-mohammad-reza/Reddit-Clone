package category

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/subreddit-domain/category/dtos"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
	"github.com/reddit-clone/src/share/services"
	"github.com/redis/go-redis/v9"
)

var ctx *gin.Context

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

func (s *CategoryService) Create(dto dtos.CreateCategoryDto, categoryId string) error {

	fmt.Println(dto)

	ParentCategory := s.pgRepository.QueryRow("SELECT * FROM category WHERE id = $1", categoryId)
	err := ParentCategory.Scan(&categoryId)
	if err != nil {
		return err
	}
	_, err = s.pgRepository.Exec("INSERT INTO category (test.categoryType, name,ParentCategory) VALUES ($1, $2, $3 )", dto.CategoryType, dto.Name, ParentCategory)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())

	}

	return nil
}
