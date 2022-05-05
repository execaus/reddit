package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/jmoiron/sqlx"
	"reddit/models"
	"time"
)

const expireSession = time.Minute

type SessionPostgres struct {
	db *sqlx.DB
}

func (r *SessionPostgres) GetAccount(hash string) (*models.Account, error) {
	var account models.Account

	if err := r.db.Get(&account, `select accountResult.* from (select * from "Session" where hash=$1
    and dead_date > current_timestamp) as sessionResult
    inner join (select * from "Account" where dead_date is null) as accountResult
	on sessionResult.account=accountResult.login`, hash); err != nil {
		return nil, err
	}

	return &account, nil
}

func NewSessionPostgres(db *sqlx.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (r *SessionPostgres) Generate(login string) (string, error) {
	timeNow := time.Now().UTC()
	sessionHash := getSessionHash(timeNow)
	timeDead := timeNow.Add(expireSession)

	_, err := r.db.Query(`insert into "Session" (hash, account, create_date, dead_date) 
					  values ($1, $2, $3, $4)`, sessionHash, login, timeNow, timeDead)
	if err != nil {
		return "", err
	}

	return sessionHash, nil
}

func getSessionHash(time time.Time) string {
	hash := getSha256Hash(time)
	return hash
}

func getSha256Hash(time time.Time) string {
	hash := sha256.New()
	hash.Write([]byte(time.String()))
	return hex.EncodeToString(hash.Sum(nil))
}
