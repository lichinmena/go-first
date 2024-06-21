package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json;"id"`
	Title  string  `json;"id"`
	Artist string  `json;"artist"`
	Price  float64 `json;"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 100.00},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 150.00},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarag Bright", Price: 200.00},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func main() {
	fmt.Println("Bienvenido a vinyl api")
	router := gin.Default()

	/*
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hola Mundo",
			})
		})
	*/
	router.GET("/albums", getAlbums)

	router.POST("/albums", addAlbum)

	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}
