package controllers

import (
	"log"
	database "my-app/Database"
	"my-app/constants"
	"my-app/models"
	"my-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductList(c *gin.Context) {
	// Get list product
	rows, err := database.DB.Query(constants.GetListProductQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	products := []models.Product{}

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Thumbnail, &product.Description, &product.CreatedAt, &product.UpdatedAt, &product.CategoryID)
		if err != nil {

			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    gin.H{"products": products},
		"message": "success!",
	})
}

func CreateProduct(c *gin.Context) {
	var newProduct models.ProductBody
	if err := c.ShouldBind(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.AreRequiredFieldsEmpty(newProduct.Name, newProduct.Price, newProduct.Description, newProduct.CategoryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "One or more required fields are empty"})
		return
	}

	uploadedURLs := utils.UploadFile(c)

	if uploadedURLs == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	newProduct.Thumbnail = uploadedURLs

	// Thực hiện truy vấn INSERT để lưu trữ thông tin sản phẩm
	_, err := database.DB.Exec(constants.CreateProductQuery, newProduct.Name, newProduct.Price, newProduct.Thumbnail, newProduct.Description, newProduct.CategoryID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully",
		"data": newProduct,
	})
}
