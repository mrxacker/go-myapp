package postgres

import (
	"database/sql"

	"github.com/mrxacker/go-myapp/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.QueryRow("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id models.UserID) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
