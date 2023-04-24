package main

import (
	"go-giner/controllers"
	"go-giner/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello go-gin"})
	})

	public := router.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:8088")
}
