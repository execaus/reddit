package service

import (
	"reddit/models"
	"reddit/pkg/repository"
)

type Post interface {
	GetById(id string) (*models.Post, error)
	GetList(page int, limit int) (*models.OutputPostList, error)
	Create(post *models.InputPost) (*models.OutputPost, error)
	Update(post *models.InputUpdatePost) error
	Delete(id string) error
}

type Auth interface {
	SignIn(input *models.InputSignIn) (*models.OutputSignIn, error)
	SignUp(input *models.InputSignUp) (*models.OutputSignUp, error)
}

type Service struct {
	Post Post
	Auth Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post: NewPostService(repos.Post),
		Auth: NewAuthService(repos.Auth),
	}
}
