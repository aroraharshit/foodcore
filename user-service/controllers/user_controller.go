package controllers

import (
	"context"

	"github.com/aroraharshit/go-foodcore/user-service/models"
	proto "github.com/aroraharshit/go-foodcore/user-service/proto"
	"github.com/aroraharshit/go-foodcore/user-service/services"
	"github.com/aroraharshit/go-foodcore/user-service/utils"
	"google.golang.org/grpc/codes"
)

type UserController struct {
	proto.UnimplementedUserServiceServer
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	registerReq := &models.RegisterUserRequest{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
	}

	resp, err := uc.service.RegisterUser(ctx, registerReq)
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{
		UserId:  resp.UserId,
		Message: resp.Message,
	}, nil
}

func (uc *UserController) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {
	loginReq := &models.LoginUserRequest{
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
	}
	
	resp, err := uc.service.LoginUser(ctx, loginReq)
	if err != nil {
		// Wrap and return a structured gRPC error
		return nil, utils.ErrorHandler(err, codes.Unauthenticated)
	}

	return &proto.AuthResponse{
		Message: resp.Message,
		UserId:  resp.UserId.Hex(),
	}, nil
}
