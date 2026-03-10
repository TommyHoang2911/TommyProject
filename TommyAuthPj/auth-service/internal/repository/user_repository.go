package repository

import (
	"database/sql"
	"time"

	"auth-service/internal/model"
)

// UserRepository performs CRUD operations on the users table.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository returns a new instance bound to the provided database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user record and updates the user object with its
// auto-generated ID.
func (r *UserRepository) Create(user *model.User) error {
	query := `
INSERT INTO users (email, password, created_at)
VALUES ($1, $2, $3)
RETURNING id
`
	user.CreatedAt = time.Now()
	return r.db.QueryRow(query, user.Email, user.Password, user.CreatedAt).Scan(&user.ID)
}
