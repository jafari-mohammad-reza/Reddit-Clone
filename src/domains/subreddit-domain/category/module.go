package category

import (
	"github.com/reddit-clone/src/api"
	"github.com/reddit-clone/src/share/config"
)

type CategoryModule struct {
	cfg        *config.Config
	controller *CategoryController
	service    *CategoryService
}

func initRoutes(c *CategoryController) {
	router := api.GetApiRoute()
	authGroup := router.Group("/category")
	authGroup.POST("/create", c.CreateCategory)
}
func NewCategoryModule() *CategoryModule {
	cfg := config.GetConfig()
	service := NewCategoryService(cfg)
	controller := NewCategoryController()
	initRoutes(controller)
	return &CategoryModule{
		cfg,
		controller,
		service,
	}
}
