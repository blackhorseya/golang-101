package main

import (
	"context"
	"log"

	"github.com/blackhorseya/golang-101/pkg/otelx"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	name       = "app1"
	addr       = ":8080"
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
	c.JSON(200, gin.H{
		"message": "Hello from app1!",
	})
}
