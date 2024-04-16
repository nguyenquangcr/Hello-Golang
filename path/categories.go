package path

import (
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCategoriesRoutes(userGroup *gin.RouterGroup) {
	userGroup.GET("/", controllers.GetCategoriesList)
	userGroup.POST("/", controllers.CreateCategory)
	userGroup.GET("/:id", controllers.GetDetailCategory)
	userGroup.PUT("/:id", controllers.UpdateCategory)
	userGroup.DELETE("/:id", controllers.DeleteCategory)
}
