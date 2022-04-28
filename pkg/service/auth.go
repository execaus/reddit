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

func (s *AuthService) SignIn(input *models.InputSignIn) (*models.OutputSignIn, error) {
	// todo send email sign-in account
	return s.repo.SignIn(input)
}

func (s *AuthService) SignUp(input *models.InputSignUp) (*models.OutputSignUp, error) {
	// todo send email registration
	return s.repo.SignUp(input)
}
