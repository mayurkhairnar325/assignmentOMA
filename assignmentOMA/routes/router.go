package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mayurkhairnar2525/restaurantManagement/controller"
	"github.com/mayurkhairnar2525/restaurantManagement/middleware"
)

func OrderRoutes(IncomingRoutes *gin.Engine) {
	orderRepo := controller.New()
	IncomingRoutes.Use(middleware.Authentication())

	IncomingRoutes.GET("/orders", orderRepo.GetAllOrders)
	IncomingRoutes.GET("/orders/:order_id", orderRepo.GetOrderByOrderID)
	IncomingRoutes.DELETE("/orders/:id", orderRepo.DeleteOrder)
	IncomingRoutes.PUT("/orders/:id", orderRepo.ModifyOrder)
	IncomingRoutes.POST("/order/orders", orderRepo.CreateOrderUsingUserID)
}
