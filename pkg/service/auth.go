package service

import (
	"github.com/golang-jwt/jwt"
	"reddit/models"
	"reddit/pkg/repository"
	"time"
)

const jwtSecret = "69738360beaec81e9c06bd7f785a1a8615522a8db1f50bc0d8ba3438880fc2cf"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignIn(input *models.InputSignIn) (*models.OutputSignIn, error) {
	// todo send email sign-in account
	output, err := s.repo.SignIn(input)
	if err != nil {
		return nil, err
	}

	token, err := generateJwtToken(output.Account.Permissions)
	output.Token = token

	return output, nil
}

func generateJwtToken(permissions string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		models.JwtClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			},
			Permissions: permissions,
		})
	signingToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signingToken, nil
}

func (s *AuthService) SignUp(input *models.InputSignUp) (*models.OutputSignUp, error) {
	// todo send email registration
	return s.repo.SignUp(input)
}
