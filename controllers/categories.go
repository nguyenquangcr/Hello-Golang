package controllers

import (
	"database/sql"
	"fmt"
	"log"
	database "my-app/Database"
	"my-app/models"
	"my-app/utils"
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
		"data":    gin.H{"categories": categories},
		"message": "success!",
	})
}

func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.IsRequiredFieldEmpty(newCategory.Name, "name") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	// Thực hiện truy vấn INSERT để lưu trữ thông tin danh mục
	_, err := database.DB.Exec("INSERT INTO categories (name) VALUES (?)", newCategory.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully",
		"data": newCategory,
	})
}

func UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var updatedCategory models.Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.IsRequiredFieldEmpty(updatedCategory.Name, "name") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	// Thực hiện truy vấn UPDATE để cập nhật thông tin danh mục
	_, err := database.DB.Exec("UPDATE categories SET name = ? WHERE id = ?", updatedCategory.Name, categoryID)
	if err != nil {
		fmt.Println("err ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully",
		"data": updatedCategory,
	})
}

func DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Thực hiện truy vấn DELETE để xóa danh mục
	_, err := database.DB.Exec("DELETE FROM categories WHERE id = ?", categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func GetDetailCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Thực hiện truy vấn SELECT để lấy thông tin chi tiết danh mục
	var category models.Category
	err := database.DB.QueryRow("SELECT * FROM categories WHERE id = ?", categoryID).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve category"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}
