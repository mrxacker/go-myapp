package service

import (
	"github.com/mrxacker/go-myapp/internal/models"
	"github.com/mrxacker/go-myapp/internal/repository"
	"github.com/mrxacker/go-myapp/pkg/logger"
)

type UserService struct {
	repo   repository.UserRepository
	logger logger.Logger
}

func NewUserService(repo repository.UserRepository, logger logger.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) GetUser(id models.UserID) (*models.User, error) {
	s.logger.Debugf("Service: getting user, %v", id)
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(username, email string) (*models.User, error) {
	s.logger.Debugf("Service: creating user, %v, %v", username, email)
	user := &models.User{
		Username: username,
		Email:    email,
	}
	return s.repo.CreateUser(user)
}
