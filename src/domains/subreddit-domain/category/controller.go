package category

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/subreddit-domain/category/dtos"
	"github.com/reddit-clone/src/share/config"
)

var ctx *gin.Context

type CategoryController struct {
	service *CategoryService
}

func NewCategoryController() *CategoryController {
	cfg := config.GetConfig()
	categoryService := NewCategoryService(cfg)
	return &CategoryController{categoryService}
}

func (c *CategoryController) CreateCategory() gin.HandlerFunc {
	dto := new(dtos.CreateCategoryDto)
	parentCategoryIdString := ctx.Params["id"]

	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err := c.service.Create(dto, parentCategoryIdString)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Created": fmt.Sprintf("Category Created With ParentId", parentCategoryId)})

}
