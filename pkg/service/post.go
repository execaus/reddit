package service

import (
	"reddit/models"
	"reddit/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetById(id string) (*models.Post, error) {
	return s.repo.GetById(id)
}
func (s *PostService) GetList(page int, limit int) (*models.OutputPostList, error) {
	return s.repo.GetList(page, limit)
}
func (s *PostService) Create(post *models.InputPost) (*models.OutputPost, error) {
	return s.repo.Create(post)
}
func (s *PostService) Update(post *models.InputUpdatePost) error {
	return s.repo.Update(post)
}
func (s *PostService) Delete(id string) error {
	return s.repo.Delete(id)
}
