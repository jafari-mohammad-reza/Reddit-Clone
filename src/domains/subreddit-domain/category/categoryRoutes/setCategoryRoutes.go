package category

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/subreddit-domain/category"
)

func CategoryRoutes(group *gin.RouterGroup) {
	categoryRoute := group.Group("category")
	categoryHandler := category.NewCategoryController()
	categoryRoute.POST("/create", categoryHandler.CreateCategory)
	return
}
