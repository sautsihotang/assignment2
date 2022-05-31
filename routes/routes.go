package routes

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartingServer() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrder)
	router.GET("/orders/:orderID", controllers.GetOrderById)
	router.DELETE("orders/:orderID", controllers.DeleteOrderById)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("/orders", controllers.DeleteAllOrder)

	return router
}
