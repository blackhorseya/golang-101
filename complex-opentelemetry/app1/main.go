package main

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/blackhorseya/golang-101/pkg/otelx"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
	err := call(c.Request.Context(), "http://localhost:8081")
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = call(c.Request.Context(), "http://localhost:8082")
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Hello from app1!",
	})
}

func call(ctx context.Context, target string) error {
	ep, err := url.ParseRequestURI(target + "/work")
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep.String(), nil)
	if err != nil {
		return err
	}

	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
