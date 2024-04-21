package path

import (
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductsRoutes(userGroup *gin.RouterGroup) {
	userGroup.GET("/", controllers.GetProductList)
	userGroup.GET("/:id", controllers.GetDetailProduct)
	userGroup.POST("/", controllers.CreateProduct)
	userGroup.PUT("/:id", controllers.UpdateProduct)
	userGroup.DELETE("/:id", controllers.DeleteProduct)
}
