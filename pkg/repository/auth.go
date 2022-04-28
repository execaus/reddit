package repository

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"reddit/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (r *AuthPostgres) SignIn(input *models.InputSignIn) (*models.OutputSignIn, error) {

}

func (r *AuthPostgres) SignUp(input *models.InputSignUp) (*models.OutputSignUp, error) {
	var output models.OutputSignUp

	passwordHash, err := getPasswordHash(input.Password)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Query(`insert into "Account" (login, password, name, email) 
		values ($1, $2, $3, $4)`, input.Login, passwordHash, input.Name, input.Email)
	if err != nil {
		return nil, err
	}

}

func getPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
