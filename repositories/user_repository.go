package repositories

import (
	"collaborative-task/db"
	"collaborative-task/models"
	"context"
	"log"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	// prepare the sql statement
	query := "INSERT INTO m_user (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"

	// Execute the statement with the pool
	row := db.Pool.QueryRow(ctx, query, user.Name, user.Username, user.Email, user.Password)

	// Scan the returned ID
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	rows, err := db.Pool.Query(ctx, "SELECT id, name, username, email, created_at, updated_at FROM m_user")
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, name, username, email, password, created_at, updated_at
		FROM m_user
		WHERE username = $1 OR email = $2
	`
	row := db.Pool.QueryRow(ctx, query, username, username)

	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return user, nil
}
