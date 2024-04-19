package path

import (
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductsRoutes(userGroup *gin.RouterGroup) {
	userGroup.GET("/", controllers.GetProductList)
	userGroup.POST("/", controllers.CreateProduct)
	// userGroup.GET("/:id", controllers.GetDetailCategory)
	// userGroup.PUT("/:id", controllers.UpdateCategory)
	// userGroup.DELETE("/:id", controllers.DeleteCategory)
}
