package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/blackhorseya/golang-101/pkg/stringx"
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

// ShortenPayload is the payload for the shorten endpoint.
type ShortenPayload struct {
	URL    string        `json:"url" binding:"required"`
	Expiry time.Duration `json:"expiry" default:"1h"`
}

func shorten(c *gin.Context) {
	var payload ShortenPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := stringx.EncodeBase62(rand.Int()) //nolint:gosec // This is just a demo.

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
