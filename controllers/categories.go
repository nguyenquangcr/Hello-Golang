package controllers

import (
	"log"
	database "my-app/Database"
	"my-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCategoriesList(c *gin.Context) {

	// Thực hiện truy vấn để lấy danh sách người dùng
	rows, err := database.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Tạo một slice để lưu trữ danh sách người dùng
	categories := []models.Category{}

	// Đọc dữ liệu từ kết quả truy vấn và lưu vào slice
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Trả về danh sách người dùng dưới dạng JSON
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
