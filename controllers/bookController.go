package controllers

import (
	"Challenge7/configs"
	"Challenge7/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var Books []models.Book

func GetAllBooks(c *gin.Context) {

	configs.GetDB().Find(&Books)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Books retrieved successfully",
		"data":    Books,
	})

}

func GetBookByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := configs.GetDB().First(&book, id)

	if res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book retrieved successfully",
		"data":    book,
	})
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	configs.GetDB().Create(&book)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Book created successfully",
		"data":    book,
	})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := configs.GetDB().First(&book, id)

	if res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	configs.GetDB().Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book updated successfully",
		"data":    book,
	})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := configs.GetDB().First(&book, id)

	if res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book not found",
		})
		return
	}

	configs.GetDB().Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book deleted successfully",
	})
}
