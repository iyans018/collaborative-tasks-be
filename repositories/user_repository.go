package repositories

import (
	"collaborative-task/db"
	"collaborative-task/models"
	"context"
	"log"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(ctx context.Context, user models.User) error {
	_, err := db.Pool.Exec(ctx, "INSERT INTO m_user (id, name, username, email, password) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error creating user:", err)
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
