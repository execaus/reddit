package service

import (
	"reddit/models"
	"reddit/pkg/repository"
)

type Post interface {
	GetById(id string) (models.Post, error)
	GetList(page int, limit int) (models.OutputPostList, error)
	Create(post models.InputPost) (models.OutputPost, error)
	Update(post models.InputUpdatePost) error
	Delete(id string) error
}

type Service struct {
	Post Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post: NewPostService(repos.Post),
	}
}
