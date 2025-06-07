package controllers

import (
	"context"
	"net/http"

	proto "github.com/aroraharshit/foodcore/api-gateway/proto"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	GRPCClient proto.UserServiceClient
}

func NewUserController(client proto.UserServiceClient) *UserController {
	return &UserController{GRPCClient: client}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Mobile   string `json:"mobile"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := uc.GRPCClient.RegisterUser(context.Background(), &proto.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":  resp.UserId,
		"message": resp.Message,
	})
}
