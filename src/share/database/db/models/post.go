package models

import (
	"time"
)

type Post struct {
	Id     int
	Title  string
	Body   string
	User   User
	UserId int
	Subreddit
	SubredditId int
	CreatedAt   time.Time
	ViewCount   int
}

type ReactionType string

const (
	Like    ReactionType = "like"
	DisLike ReactionType = "dislike"
)

type PostReaction struct {
	Id int
	Post
	PostId int
	User
	UserId    int
	Reaction  ReactionType
	CreatedAt time.Time
}
