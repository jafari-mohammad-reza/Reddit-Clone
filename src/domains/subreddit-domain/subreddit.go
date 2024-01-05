package subreddit_domain

import "github.com/reddit-clone/src/domains/subreddit-domain/category"

type SubredditDomain struct {
	categoryModule *category.CategoryModule
}

func NewSubredditDomain() *SubredditDomain {
	categoryModule := category.NewCategoryModule()
	return &SubredditDomain{
		categoryModule,
	}
}
