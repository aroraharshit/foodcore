package routes

import (
	"github.com/aroraharshit/foodcore/api-gateway/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userController *controllers.UserController) {
	r.POST("/register", userController.RegisterUser)
	r.POST("/login",userController.LoginUser)
}
