package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"reddit/models"
	"strings"
	"time"
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
	id := uuid.New().String()
	if id == "" {
		return nil, errors.New("generate uuid invalid")
	}

	timeNow := time.Now()
	_, err := r.db.Query(`insert into "Post" (id, author, caption, body, create_date, deleted) 
						values ($1, $2, $3, $4, $5, $6)`,
		id, post.Author, post.Caption, post.Body, timeNow, false)
	if err != nil {
		return nil, err
	}

	return &models.OutputPost{
		Id:         id,
		CreateDate: timeNow,
	}, nil
}
func (r *PostPostgres) Update(post *models.InputUpdatePost) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if post.Caption != "" {
		setValues = append(setValues, fmt.Sprintf("caption=$%d", argId))
		args = append(args, post.Caption)
		argId++
	}

	if post.Body != "" {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argId))
		args = append(args, post.Body)
		argId++
	}

	querySetPart := strings.Join(setValues, ", ")
	args = append(args, post.Id)

	query := fmt.Sprintf(`update "Post" set %s where id=$%d`, querySetPart, argId)

	_, err := r.db.Query(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostPostgres) Delete(id string) error {
	_, err := r.db.Query(`update "Post" set deleted=true where id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}
