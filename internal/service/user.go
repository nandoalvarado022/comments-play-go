package service

import (
	"errors"

	"github.com/nandoalvarado022/comments-play/internal/models"
	"github.com/nandoalvarado022/comments-play/internal/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(r repository.Repository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByUID(uid string) (*models.User, error) {
	if uid == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	user, err := s.repo.GetByUID(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
