package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/subreddit-domain/category"
	"github.com/reddit-clone/src/domains/subreddit-domain/category/dtos"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/services"
	"github.com/redis/go-redis/v9"
)

type CategoryService struct {
	redisClient     *redis.Client
	rabbitMqService *services.RabbitMQService
	pgRepository    *sql.DB
}

func NewCategoryService() *CategoryService {
	pg := postgres.GetPostgres()
	return &CategoryService{
		redisClient:     cache.GetRedisClient(),
		pgRepository:    pg,
		rabbitMqService: services.NewRabbitMQService(lg, "category", nil), // this setting just for now
	}
}
func CreateHandler(ctx *gin.Context) error {
	categoryService := category.NewCategoryService()

	asdasd := category.NewCategoryService()
	var dto dtos.CreateCategoryDto
	time.Sleep(77 * time.Second)
	test := new(dtos.CreateCategoryDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, test)
	}
	fmt.Println(test)

	_, err := categoryService.pgRepository.Exec("INSERT INTO category (test.categoryType, name,ParentCategory) VALUES ($1, $2 )", test.CategoryType, test.Name)

	if err != nil {
		return nil, err
	}

	return nil
}
