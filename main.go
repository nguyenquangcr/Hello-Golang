package main

import (
	database "my-app/Database"
	"my-app/constants"
	"my-app/path"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Func connection databse
	database.InitDB(constants.ConnectionString)

	router := gin.Default()
	api := router.Group("/api/v1")
	{

		// Router Group
		user := api.Group("/user")
		{
			path.SetupUserRoutes(user)
		}
		categories := api.Group("/categories")
		{
			path.SetupCategoriesRoutes(categories)
		}
		products := api.Group("/products")
		{
			path.SetupProductsRoutes(products)
		}
	}

	router.Run(":8080")
}
