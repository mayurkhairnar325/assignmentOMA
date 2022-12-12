package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mayurkhairnar2525/restaurantManagement/controller"
	"github.com/mayurkhairnar2525/restaurantManagement/middleware"
)

func UserRoutes(IncomingRoutes *gin.Engine) {
	UserRepo := controller.New()
	IncomingRoutes.POST("/users/signup", UserRepo.SignUp)
	IncomingRoutes.POST("/users/login", UserRepo.Login)

	IncomingRoutes.Use(middleware.Authentication())
	IncomingRoutes.GET("/users", UserRepo.GetUsers)
	IncomingRoutes.GET("/user/userid/:id", UserRepo.GetUserByID)
	IncomingRoutes.GET("/user/use/:user_id", UserRepo.GetUserBy_userID)

}
