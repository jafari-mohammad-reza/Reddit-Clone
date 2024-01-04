package models

import (
	"time"
)

type User struct {
	Id                int
	Username          string `json:"username"`
	Email             string `json:"email"` // validate:"required,email"
	Password          string `json:"password"`
	Bio               string `json:"bio"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
}

type UserFollower struct {
	Id         int
	UserId     int
	FollowerId int
	FollowedAt time.Time
}
