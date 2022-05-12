package repository

import (
	"github.com/jmoiron/sqlx"
	"reddit/models"
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
	Generate(login string) (string, error)
	GetAccount(hash string) (*models.Account, error)
}

type RecoverAccess interface {
	GenerateLink(input *models.InputRecoverAccessLink) (string, string, error)
	RegisterNewPassword(input *models.InputRecoverAccessRegister) (string, string, error)
}

type Repository struct {
	Post          Post
	Auth          Auth
	Session       Session
	RecoverAccess RecoverAccess
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post:          NewPostPostgres(db),
		Auth:          NewAuthPostgres(db),
		Session:       NewSessionPostgres(db),
		RecoverAccess: NewRecoverAccessPostgres(db),
	}
}
