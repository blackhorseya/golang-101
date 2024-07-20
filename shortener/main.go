package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/shorten", shorten)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func shorten(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "shorten",
	})
}
