package main

import (
	database "my-app/Database"
	"my-app/path"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Kết nối tới cơ sở dữ liệu MySQL
	connStr := "root:@tcp(localhost:3306)/shopapp"
	database.InitDB(connStr)

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
	}

	router.Run(":8080")
}
