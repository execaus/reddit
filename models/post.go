package models

import "time"

type Post struct {
	Id         string    `json:"id" db:"id"`
	Author     string    `json:"author" db:"author"`
	Caption    string    `json:"caption" db:"caption"`
	Body       string    `json:"body" db:"body"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	Deleted    bool      `json:"-" db:"deleted"`
}

type OutputPostList struct {
	Posts      []Post `json:"posts"`
	TotalCount int    `json:"total_count"`
}

type InputPost struct {
	Author  string `binding:"required"`
	Caption string `binding:"required"`
	Body    string `binding:"required"`
}

type OutputPost struct {
	Id         string
	CreateDate time.Time
}

type InputUpdatePost struct {
	Id      string
	Caption string
	Body    string
}

type UriGetPostList struct {
	Page  string `uri:"page"`
	Limit string `uri:"limit"`
}
