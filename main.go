package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/handlers"
	"web-service-gin/middleware"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, value := range albums {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default() // Initialize Gin using Default
	//router.GET("/albums", getAlbum)
	//router.GET("/albums/:id", getAlbumByID)
	//router.POST("/albums", postAlbum)

	// Public routes (do not require authentication)
	publicRoutes := router.Group("/public") {
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/login", handlers.Register)
	}
	// Protected routes (requires authentication)
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		// Protected routes here
	}

	router.Run("localhost:8080")
}
