package repository

import (
	"github.com/jmoiron/sqlx"
	"reddit/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (a *AuthPostgres) SignIn(input *models.InputSignIn) (*models.OutputSignIn, error) {

}

func (a *AuthPostgres) SignUp(input *models.InputSignUp) (*models.OutputSignUp, error) {

}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
