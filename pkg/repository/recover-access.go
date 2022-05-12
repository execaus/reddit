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

func (r *RecoverAccessPostgres) RegisterNewPassword(input *models.InputRecoverAccessRegister) (string, string, error) {
	var account models.Account
	tx, err := r.db.Beginx()
	if err != nil {
		return "", "", err
	}

	if err = tx.Get(&account, `select acc.* from (select * from "RecoverAccess" where id=$1 
        and dead_date>current_timestamp and completed=false) as ra
    	inner join (select * from "Account") as acc on ra.account=acc.login`, input.Id); err != nil {
		_ = tx.Rollback()
		return "", "", err
	}

	passwordHash, err := getPasswordHash(input.Password)
	if err != nil {
		_ = tx.Rollback()
		return "", "", err
	}

	_, err = tx.Query(`update "Account" set password=$1 where login=$2`, passwordHash, account.Login)
	if err != nil {
		_ = tx.Rollback()
		return "", "", err
	}

	_, err = tx.Query(`update "RecoverAccess" set completed=true where id=$1`, input.Id)
	if err != nil {
		_ = tx.Rollback()
		return "", "", err
	}

	if err = tx.Commit(); err != nil {
		return "", "", err
	}

	return account.Email, "", nil
}
