package path

import (
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCategoriesRoutes(userGroup *gin.RouterGroup) {
	userGroup.GET("/", controllers.GetCategoriesList)
	// userGroup.GET("/:id", controllers.GetUserByID)
	// userGroup.POST("/", controllers.CreateUser)
	// userGroup.PUT("/:id", controllers.UpdateUser)
	// userGroup.DELETE("/:id", controllers.DeleteUser)
}
