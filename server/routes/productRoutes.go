package routes

import (
	controller "golang-chain-management/controllers"
	"golang-chain-management/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	auth := middleware.Authentication()

	// Use the authenticated group for the product routes
	productGroup := incomingRoutes.Group("/product")
	{
		// Apply authentication middleware to each route in the product group
		productGroup.Use(auth)

		productGroup.GET("", controller.GetProducts())
		productGroup.GET("/:product_id", controller.GetProduct())
		productGroup.POST("/create", controller.CreateProduct())
		productGroup.POST("/update", controller.UpdateProduct())
		productGroup.DELETE("/delete", controller.DeleteProduct())
	}
}
