package subreddit_domain

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/subreddit-domain/category"
	"github.com/reddit-clone/src/domains/subreddit-domain/subreddit"
)

type SubredditDomain struct {
	categoryModule *category.CategoryModule
	subredditModule *subreddit.SubredditModule
}

func NewSubredditDomain(r *gin.RouterGroup) *SubredditDomain {
	categoryModule := category.NewCategoryModule(r)
	subredditModule := subreddit.NewSubredditModule(r)
	return &SubredditDomain{
		categoryModule,
		subredditModule,
	}
}
