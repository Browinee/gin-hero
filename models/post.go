package models

import "time"

type Post struct {
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	AuthorId    uint64    `json:"author_id" db:"author_id"`
	PostID      uint64    `json:"post_id" db:"post_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}
