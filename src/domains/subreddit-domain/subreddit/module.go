package subreddit

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/share/config"
)

type SubredditModule struct {
	cfg        *config.Config
	controller *SubredditController
	service    *SubredditService
}

func initRoutes(r *gin.RouterGroup , c *SubredditController) {
	authGroup := r.Group("/subreddit")
	authGroup.GET("/", c.Subreddits)
}
func NewSubredditModule(r *gin.RouterGroup ) *SubredditModule {
	cfg := config.GetConfig()
	service := NewSubredditService(cfg)
	controller := NewSubredditController(service)
	initRoutes(r , controller)
	return &SubredditModule{
		cfg,
		controller,
		service,
	}
}
