package repository

import (
	"github.com/jmoiron/sqlx"
	"reddit/models"
)

type Post interface {
	GetById(id string) (models.Post, error)
	GetList(page int, limit int) (models.OutputPostList, error)
	Create(post models.InputPost) (models.OutputPost, error)
	Update(post models.InputUpdatePost) error
	Delete(id string) error
}

type Repository struct {
	Post Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: NewPostPostgres(db),
	}
}
