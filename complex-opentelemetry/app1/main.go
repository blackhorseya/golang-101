package main

import (
	"log"
	"net/http"
	"time"

	"github.com/blackhorseya/golang-101/pkg/timex"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const name = "app1"
const addr = ":8080"

func main() {
	router := gin.Default()
	router.Use(otelgin.Middleware(name))
	router.GET("/work", work)

	err := router.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func work(c *gin.Context) {
	duration := timex.GetRandomDuration(1, 3)
	time.Sleep(duration)

	c.JSON(http.StatusOK, gin.H{
		"message":         "work done",
		"duration_second": duration.Seconds(),
	})
}
