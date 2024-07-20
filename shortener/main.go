package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/blackhorseya/golang-101/pkg/storage/redix"
	"github.com/blackhorseya/golang-101/pkg/stringx"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	client, err := redix.NewClient(redix.Options{Addr: "localhost:6379"})
	if err != nil {
		panic(err)
	}
	rdb = client

	router := gin.Default()
	router.POST("/shorten", shorten)
	router.GET("/s/:id", resolve)

	err = router.Run(":8080")
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

	err = rdb.Set(context.Background(), id, payload.URL, payload.Expiry).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func resolve(c *gin.Context) {
	id := c.Param("id")

	url, err := rdb.Get(context.Background(), id).Result()
	if errors.Is(err, redis.Nil) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
