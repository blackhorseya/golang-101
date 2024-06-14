package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/blackhorseya/golang-101/pkg/otelx"
	"github.com/blackhorseya/golang-101/pkg/timex"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	name       = "app2"
	addr       = ":8081"
	otelTarget = "localhost:4317"
)

func main() {
	// Setup OpenTelemetry SDK
	err := otelx.SetupOTelSDK(context.Background(), otelTarget, name)
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry SDK: %v", err)
	}
	defer func() {
		err = otelx.Shutdown(context.Background())
		if err != nil {
			log.Printf("Failed to shutdown OpenTelemetry SDK: %v", err)
		}
	}()

	// Setup Gin router
	router := gin.Default()
	router.Use(otelgin.Middleware(name))
	router.GET("/work", work)

	// Start the server
	err = router.Run(addr)
	if err != nil {
		log.Panicf("Failed to start the server: %v", err)
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
