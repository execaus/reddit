package models

import "time"

type Post struct {
	Id         string
	Author     string
	Caption    string
	Body       string
	CreateDate time.Time
}

type OutputPostList struct {
	Posts      []Post
	TotalCount int
}

type InputPost struct {
	Author  string
	Caption string
	Body    string
}

type OutputPost struct {
	Id         string
	CreateDate time.Time
}

type InputUpdatePost struct {
	Caption string
	Body    string
}
