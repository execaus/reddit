package service

import (
	"reddit/models"
	"reddit/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) SignIn(input *models.InputSignIn) (*models.OutputSignIn, error) {

}

func (a *AuthService) SignUp(input *models.InputSignUp) (*models.OutputSignUp, error) {

}
