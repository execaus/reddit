package repository

import (
	"github.com/jmoiron/sqlx"
	"reddit/models"
	"time"
)

const expireRecoverAccessLink = time.Minute * 10

type RecoverAccessPostgres struct {
	db *sqlx.DB
}

func NewRecoverAccessPostgres(db *sqlx.DB) *RecoverAccessPostgres {
	return &RecoverAccessPostgres{db: db}
}

func (r *RecoverAccessPostgres) GenerateLink(input *models.InputRecoverAccessLink) (string, string, error) {
	var account models.Account

	if err := r.db.Get(&account,
		`select * from "Account" where login=$1 or email=$1 and dead_date is null`, input.Account); err != nil {
		return "", "", err
	}

	timeNow := time.Now().UTC()
	link := getSha256Hash(timeNow)

	_, err := r.db.Query(`insert into "RecoverAccess" (id, account, dead_date, completed) values ($1, $2, $3, $4)`,
		link, account.Login, timeNow.Add(expireRecoverAccessLink), false)
	if err != nil {
		return "", "", err
	}

	return link, account.Email, nil
}
