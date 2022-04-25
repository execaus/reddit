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

func (r *PostPostgres) GetById(id string) (*models.Post, error) {
	var post models.Post

	if err := r.db.Get(&post, `select * from "Post" where id=$1`, id); err != nil {
		return nil, err
	}

	return &post, nil
}
func (r *PostPostgres) GetList(page int, limit int) (*models.OutputPostList, error) {
	var output models.OutputPostList

	if err := r.db.Select(&output.Posts, `select * from "Post" where deleted=false 
                     order by create_date desc limit $1 offset $2`,
		limit, (page-1)*limit); err != nil {
		return nil, err
	}

	if err := r.db.Get(&output.TotalCount, `select count(*) from "Post" where deleted=false`); err != nil {
		return nil, err
	}

	return &output, nil
}
func (r *PostPostgres) Create(post *models.InputPost) (*models.OutputPost, error) {

}
func (r *PostPostgres) Update(post *models.InputUpdatePost) error {

}
func (r *PostPostgres) Delete(id string) error {

}
