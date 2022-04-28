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
	SignUp(input *models.InputSignUp) (*models.OutputSignUp, error)
}

type Session interface {
	Generate(login string) (string, error)
}

type Repository struct {
	Post Post
	Auth Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: NewPostPostgres(db),
		Auth: NewAuthPostgres(db),
	}
}
