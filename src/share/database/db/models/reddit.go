package models

import "time"

type Subreddit struct {
	Id          int
	name        string
	Description string
	Owner       User
	OwnerId     int
	Category    string
	CreatedAt   time.Time
}
