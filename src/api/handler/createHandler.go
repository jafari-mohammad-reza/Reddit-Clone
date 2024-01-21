package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/subreddit-domain/category/dtos"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
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

func NewCategoryService() *CategoryService {
	pg := postgres.GetPostgres()
	return &CategoryService{
		redisClient:  cache.GetRedisClient(),
		pgRepository: pg,
	}
}
func CreateHandler() gin.HandlerFunc {

	asdasd := NewCategoryService()
	test := new(dtos.CreateCategoryDto)
	if err := ctx.ShouldBindJSON(test); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(test)

	_, err := asdasd.pgRepository.Exec("INSERT INTO category (test.categoryType, name,ParentCategory) VALUES ($1, $2 )", test.CategoryType, test.Name)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())

	}
	ctx.JSON(http.StatusAccepted, &test)
	return nil
}
