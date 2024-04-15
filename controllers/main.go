package controllers

import (
	"my-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world!",
	})
}

func GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List user nè!",
	})
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	var detailUser models.User
	if err := c.ShouldBindJSON(&detailUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin người dùng với ID tương ứng
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Detail user nè!", "userID": userID})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lưu trữ thông tin người dùng
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin người dùng với ID tương ứng
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "userID": userID})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Xóa thông tin người dùng với ID tương ứng
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "userID": userID})
}
