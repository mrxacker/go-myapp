package models

type UserID string
type User struct {
	ID       UserID `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
