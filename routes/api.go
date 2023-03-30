package routes

import (
	"Challenge7/controllers"
	"github.com/gin-gonic/gin"
)

func ApiInit() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.GET("/books/:id", controllers.GetBookByID)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
