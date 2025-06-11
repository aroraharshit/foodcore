package controllers

import (
	"context"
	"errors"

	"github.com/aroraharshit/foodcore/api-gateway/models"
	proto "github.com/aroraharshit/foodcore/api-gateway/proto"
	"github.com/aroraharshit/foodcore/api-gateway/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	GRPCClient proto.UserServiceClient
}

func NewUserController(client proto.UserServiceClient) *UserController {
	return &UserController{GRPCClient: client}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var req models.RegisterUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.ResponseHandler(c, nil, err)
		return
	}

	resp, err := uc.GRPCClient.RegisterUser(context.Background(), &proto.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		utils.ResponseHandler(c, nil, err)
		return
	}

	utils.ResponseHandler(c, gin.H{
		"userId":  resp.UserId,
		"message": resp.Message,
	}, nil)
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var req models.LoginUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.ResponseHandler(c, nil, err)
	}

	resp, err := uc.GRPCClient.LoginUser(context.Background(), &proto.LoginRequest{
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		utils.ResponseHandler(c, nil, err)
		return
	}

	if resp == nil {
		utils.ResponseHandler(c, nil, errors.New("user doesn't exist"))
		return
	}

	utils.ResponseHandler(c, gin.H{
		"userId":  resp.UserId,
		"message": resp.Message,
	}, nil)
}
