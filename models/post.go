package models

import "time"

type Post struct {
	Id          int64     `json:"id"`
	PostContent string    `json:"post_content"`
	CreatedAt   time.Time `json:"created_at"`
	UserId      int64     `json:"user_id"`
}
