package repository

import "github.com/mrxacker/go-myapp/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id models.UserID) (*models.User, error)
}
