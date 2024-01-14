package category

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/share/config"
)

type CategoryModule struct {
	cfg        *config.Config
	controller *CategoryController
	service    *CategoryService
}

func initRoutes(r *gin.RouterGroup ,  c *CategoryController) {
	authGroup := r.Group("/category")
	authGroup.GET("/", c.Categories)
}
func NewCategoryModule(r *gin.RouterGroup) *CategoryModule {
	cfg := config.GetConfig()
	service := NewCategoryService(cfg)
	controller := NewCategoryController(service)
	initRoutes(r , controller)
	return &CategoryModule{
		cfg,
		controller,
		service,
	}
}
