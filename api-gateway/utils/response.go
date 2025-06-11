package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ResponseHandler(c *gin.Context, data any, err error) {
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func handleError(c *gin.Context, err error) {
	if grpcErr, ok := status.FromError(err); ok {
		switch grpcErr.Code() {
		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": grpcErr.Message()})
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": grpcErr.Message()})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": grpcErr.Message()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": grpcErr.Message()})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
}
