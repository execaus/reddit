package service

import (
	"reddit/models"
	"reddit/pkg/repository"
)

type RecoverAccessService struct {
	repo repository.RecoverAccess
}

func NewRecoverAccessService(repo repository.RecoverAccess) *RecoverAccessService {
	return &RecoverAccessService{repo: repo}
}

func (r *RecoverAccessService) GenerateLink(input *models.InputRecoverAccessLink) error {
	link, email, err := r.repo.GenerateLink(input)
	if err != nil {
		return err
	}

	sendEmailRecoverAccess(link, email)

	return nil
}
