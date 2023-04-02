package controllers

import (
	"Challenge7/configs"
	"Challenge7/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

var Books []models.Book

func GetAllBooks(c *gin.Context) {

	res := configs.GetDB().Find(&Books)

	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Books)
}

func GetBookByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := configs.GetDB().First(&book, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	} else if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := configs.GetDB().Create(&book)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	} else if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := configs.GetDB().First(&book, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	} else if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	configs.GetDB().Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
	})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := configs.GetDB().First(&book, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	} else if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	res = configs.GetDB().Delete(&book)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
