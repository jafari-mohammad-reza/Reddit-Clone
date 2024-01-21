package category

import "github.com/gin-gonic/gin"

type CategoryController struct {
	service *CategoryService
}

func NewCategoryController(service *CategoryService) *CategoryController {
	return &CategoryController{service}
}

func (c *CategoryController) Categories(ctx *gin.Context) {
	ctx.JSON(200, "categories")
}
