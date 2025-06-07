package controllers

import (
	"context"

	"github.com/aroraharshit/go-foodcore/user-service/models"
	proto "github.com/aroraharshit/go-foodcore/user-service/proto"
	"github.com/aroraharshit/go-foodcore/user-service/services"
)

type UserController struct {
	proto.UnimplementedUserServiceServer
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	userReq := &models.RegisterUserRequest{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := uc.service.RegisterUser(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{
		UserId:  res.UserId,
		Message: res.Message,
	}, nil
}
