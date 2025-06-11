package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aroraharshit/go-foodcore/user-service/models"
	"github.com/aroraharshit/go-foodcore/user-service/repositories"
	"github.com/aroraharshit/go-foodcore/user-service/utils"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error)
	LoginUser(ctx context.Context, req *models.LoginUserRequest) (*models.LoginUserResponse, error)
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

func (us *userService) LoginUser(ctx context.Context, req *models.LoginUserRequest) (*models.LoginUserResponse, error) {
	// emailExists, err := us.repo.FindByEmail(ctx, req.Email)
	// if err != nil {
	// 	return nil, err
	// }

	// if !emailExists {
	// 	return nil, errors.New("User with this email doesnt exists")
	// }

	mobileExists, userId, err := us.repo.FindByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}

	if !mobileExists {
		return nil, errors.New("user with this mobile number doesnt exists")
	}

	storedHash, err := us.repo.FetchHashPassword(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	isPasswordMatched, err := utils.VerifyPassword(req.Password, storedHash)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !isPasswordMatched {
		return nil, errors.New("password doesnt matched")
	}

	return &models.LoginUserResponse{
		UserId:  userId,
		Message: "Successfully Login",
	}, nil
}
