package category

import (
	"Reddit-Clone/src/share/config"
	"github.com/reddit-clone/src/api"
)

type CategoryModule struct {
	cfg        *config.Config
	controller *CategoryController
	service    *CategoryService
}

func initRoutes(c *CategoryController) {
	router := api.GetApiRoute()
	authGroup := router.Group("/category")
	authGroup.GET("/", c.Categories)
}
func NewCategoryModule() *CategoryModule {
	cfg := config.GetConfig()
	service := NewCategoryService(cfg)
	controller := NewCategoryController(service)
	initRoutes(controller)
	return &CategoryModule{
		cfg,
		controller,
		service,
	}
}
