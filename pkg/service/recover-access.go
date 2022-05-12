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

func (s *RecoverAccessService) GenerateLink(input *models.InputRecoverAccessLink) error {
	link, email, err := s.repo.GenerateLink(input)
	if err != nil {
		return err
	}

	sendEmailRecoverAccess(link, email)

	return nil
}

func (s *RecoverAccessService) RegisterNewPassword(input *models.InputRecoverAccessRegister) (string, error) {
	email, message, err := s.repo.RegisterNewPassword(input)
	if err != nil {
		return message, err
	}
	sendEmail("Изменение пароля прошло успешно!", email)
	return "", nil
}
