package service

import (
	"Gophermart/internal/app/domain/model"
	"Gophermart/internal/app/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserService struct {
	userRepo *repository.UserRepository
	Logger   *zap.Logger
}

func NewUserService(userRepo *repository.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{userRepo: userRepo, Logger: logger}
}

func (s *UserService) RegisterNewUser(ctx context.Context, request *AuthRequest) (int, error) {
	eUser, err := s.userRepo.FindByLogin(ctx, model.Username(request.Login))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	if eUser != nil {
		return 0, ErrUserAlreadyExists
	}

	user, err := model.NewUserFromRealPass(request.Login, request.Password)
	if err != nil {
		return 0, fmt.Errorf("NewUserFromRealPass: %w", err)
	}
	userId, err := s.userRepo.AddNew(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("userRepo.AddNew: %w", err)
	}

	return userId, nil
}

func (s *UserService) Login(ctx context.Context, request *AuthRequest) (int, *model.User, error) {
	eUser, err := s.userRepo.FindByLogin(ctx, model.Username(request.Login))
	if err != nil {
		return 0, nil, err
	}
	user := model.NewUserFromEntity(eUser)
	err = user.CheckPassword([]byte(request.Password))
	if err != nil {
		return 0, nil, err
	}

	return eUser.ID, user, nil
}
