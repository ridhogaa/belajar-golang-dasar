package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	Name  string
	Price int
}

var albums = []Album{
	{
		"test",
		100,
	},
	{
		"test",
		100,
	},
	{
		"test",
		100,
	},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
