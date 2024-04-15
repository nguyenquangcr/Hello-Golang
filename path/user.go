package path

import (
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(userGroup *gin.RouterGroup) {
	userGroup.GET("/hello", controllers.GetHello)
	userGroup.GET("/", controllers.GetUserList)
	userGroup.GET("/:id", controllers.GetUserByID)
	userGroup.POST("/", controllers.CreateUser)
	userGroup.PUT("/:id", controllers.UpdateUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)
}
