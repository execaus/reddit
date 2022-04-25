package repository

import (
	"github.com/jmoiron/sqlx"
	"reddit/models"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) GetById(id string) (models.Post, error) {

}
func (r *PostPostgres) GetList(page int, limit int) (models.OutputPostList, error) {

}
func (r *PostPostgres) Create(post models.InputPost) (models.OutputPost, error) {

}
func (r *PostPostgres) Update(post models.InputUpdatePost) error {

}
func (r *PostPostgres) Delete(id string) error {

}
