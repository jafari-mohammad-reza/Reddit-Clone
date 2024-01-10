package subreddit_domain

import (
	"Reddit-Clone/src/domains/subreddit-domain/category"
)

type SubredditDomain struct {
	categoryModule *category.CategoryModule
}

func NewSubredditDomain() *SubredditDomain {
	categoryModule := category.NewCategoryModule()
	return &SubredditDomain{
		categoryModule,
	}
}
