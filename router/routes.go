package router

import (
	"my-gin-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func Use(route *gin.Engine) {
	route.GET("", controllers.HomeIndex)

	products := route.Group("products")
	products.GET("", controllers.GetProducts)
	products.POST("", controllers.CreateProduct)
	products.GET("/:id", controllers.GetProductByID)
	products.PUT("/:id", controllers.UpdateProduct)
	products.DELETE("/:id", controllers.DeleteProduct)

}
