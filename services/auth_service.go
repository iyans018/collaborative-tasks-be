package services

import (
	"collaborative-task/models"
	"collaborative-task/repositories"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, username string, password string) (*models.User, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	user.Password = ""

	return nil
}

func (s *authService) Login(ctx context.Context, username string, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		print(err.Error())
		return nil, errors.New("invalid username or password")
	}

	fmt.Printf("user password: %s \n", user.Password)
	fmt.Printf("input password: %s \n", password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		print(err.Error())
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
