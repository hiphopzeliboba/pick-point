package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"pickpoint/internal/model"
	"pickpoint/internal/repository"
)

type Service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, email, password string, role model.Role) (*model.User, error) {
	if !model.IsValidRole(role) {
		return nil, model.ErrInvalidRole
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	return s.repo.CreateUser(ctx, u)
}

func (s *Service) Login(ctx context.Context, email, password string) (*model.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, model.ErrInvalidCredentials
	}

	return u, nil
}
