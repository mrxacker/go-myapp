package service

import (
	"github.com/mrxacker/go-myapp/internal/models"
	"github.com/mrxacker/go-myapp/internal/repository"
	"github.com/mrxacker/go-myapp/pkg/logger"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(id models.UserID) (*models.User, error) {
	logger.Get().Debugf("Service: getting user, %v", id)
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(username, email string) (*models.User, error) {
	logger.Get().Debugf("Service: creating user, %v, %v", username, email)
	user := &models.User{
		Username: username,
		Email:    email,
	}
	return s.repo.CreateUser(user)
}
