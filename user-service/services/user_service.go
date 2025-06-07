package services

import (
	"context"
	"time"

	"github.com/aroraharshit/go-foodcore/user-service/models"
	"github.com/aroraharshit/go-foodcore/user-service/repositories"
	"github.com/aroraharshit/go-foodcore/user-service/utils"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	exists, err := us.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return &models.RegisterUserResponse{
			Message: "User already exists",
		}, nil
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.RegistertUserInsertion{
		Name:      req.Name,
		Mobile:    req.Mobile,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userId, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.RegisterUserResponse{UserId: userId, Message: "User register successfully"}, nil
}
