package repository

import "github.com/jmoiron/sqlx"

type SessionPostgres struct {
	db *sqlx.DB
}

func NewSessionPostgres(db *sqlx.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (s *SessionPostgres) Generate(login string) (string, error) {

}
