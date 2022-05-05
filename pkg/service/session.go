package service

import (
	"reddit/models"
	"reddit/pkg/repository"
)

type SessionService struct {
	repo repository.Session
}

func (s *SessionService) GetAccount(hash string) (*models.Account, error) {
	return s.repo.GetAccount(hash)
}

func NewSessionService(repo repository.Session) *SessionService {
	return &SessionService{repo: repo}
}
