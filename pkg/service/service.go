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
	SignUp(input *models.InputSignUp) error
}

type Session interface {
	GetAccount(hash string) (*models.Account, error)
	Generate(login string) (string, error)
}

type RecoverAccess interface {
	GenerateLink(link *models.InputRecoverAccessLink) error
	RegisterNewPassword(input *models.InputRecoverAccessRegister) (string, error)
}

type Service struct {
	Post          Post
	Auth          Auth
	Session       Session
	RecoverAccess RecoverAccess
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post:          NewPostService(repos.Post),
		Auth:          NewAuthService(repos.Auth),
		Session:       NewSessionService(repos.Session),
		RecoverAccess: NewRecoverAccessService(repos.RecoverAccess),
	}
}
