package controllers

import (
	"database/sql"
	"log"
	database "my-app/Database"
	"my-app/constants"
	"my-app/models"
	"my-app/utils"
	"net/http"
	"path"

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
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Thumbnail,
			&product.Description, &product.CreatedAt, &product.UpdatedAt, &product.CategoryID)
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

func GetDetailProduct(c *gin.Context) {
	productID := c.Param("id")

	var detailProduct models.Product
	err := database.DB.QueryRow(constants.GetDetailProductQuery, productID).Scan(&detailProduct.ID, &detailProduct.Name, &detailProduct.Price,
		&detailProduct.Thumbnail, &detailProduct.Description, &detailProduct.CreatedAt, &detailProduct.UpdatedAt, &detailProduct.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": detailProduct, "message": "Success"})
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

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var oldProduct models.Product
	err := database.DB.QueryRow(constants.GetDetailProductUpdateQuery, productID).Scan(&oldProduct.Name, &oldProduct.Price, &oldProduct.Thumbnail, &oldProduct.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	var newProduct models.ProductBody
	if err := c.ShouldBind(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newProduct.Name == "" {
		newProduct.Name = oldProduct.Name
	}
	if newProduct.Description == "" {
		newProduct.Description = oldProduct.Description
	}
	if newProduct.Price == 0 {
		newProduct.Price = oldProduct.Price
	}

	if len(newProduct.Files) > 0 {
		uploadedURLs := utils.ChangeFileUpload(c, oldProduct.Thumbnail)
		if uploadedURLs == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
			return
		}
		newProduct.Thumbnail = uploadedURLs
	} else {
		newProduct.Thumbnail = oldProduct.Thumbnail
	}

	// Thực hiện truy vấn UPDATE để cập nhật thông tin danh mục
	_, errUpdate := database.DB.Exec(constants.UpdateProductQuery,
		newProduct.Name, newProduct.Price, newProduct.Thumbnail, newProduct.Description, productID)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully",
		"data": newProduct,
	})
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	var thumbnail string
	errGetDetail := database.DB.QueryRow(constants.GetThumbnailQuery, productID).Scan(&thumbnail)
	if errGetDetail != nil {
		if errGetDetail == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errGetDetail})
		}
		return
	}

	errDeleteFile := utils.DeleteFile(path.Base(thumbnail))
	if errDeleteFile != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDeleteFile})
		return
	}

	_, err := database.DB.Exec(constants.DeleteProductQuery, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
