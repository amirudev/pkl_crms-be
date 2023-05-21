package routes

import (
	"github.com/amirudev/pkl_crms-be/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	v1 := api.Group("/v1")

	customer := v1.Group("/customer")
	{
		customer.GET("/", controllers.GetCustomers)
		customer.GET("/:id", controllers.GetCustomer)
		customer.POST("/", controllers.CreateCustomer)
		customer.PUT("/:id", controllers.UpdateCustomer)
		customer.DELETE("/:id", controllers.DeleteCustomer)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
