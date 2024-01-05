package subreddit

import (
	"github.com/reddit-clone/src/api"
	"github.com/reddit-clone/src/share/config"
)

type SubredditModule struct {
	cfg        *config.Config
	controller *SubredditController
	service    *SubredditService
}

func initRoutes(c *SubredditController) {
	router := api.GetApiRoute()
	authGroup := router.Group("/category")
	authGroup.GET("/", c.Categories)
}
func NewSubredditModule() *SubredditModule {
	cfg := config.GetConfig()
	service := NewSubredditService(cfg)
	controller := NewSubredditController(service)
	initRoutes(controller)
	return &SubredditModule{
		cfg,
		controller,
		service,
	}
}
