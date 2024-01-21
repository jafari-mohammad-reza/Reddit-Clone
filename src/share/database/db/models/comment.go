package models

import "time"

type Comment struct {
	ID int `json:"id"`
	Post
	PostID          int `json:"post_id"`
	ParentComment   *Comment
	ParentCommentID int `json:"parent_comment_id"`
	User
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ViewCount int       `json:"view_count"`
}
type CommentReaction struct {
	ID        int          `json:"id"`
	CommentID int          `json:"comment_id"`
	UserID    int          `json:"user_id"`
	Reaction  ReactionType `json:"reaction"`
	CreatedAt time.Time    `json:"created_at"`
}
