package main

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"

	"github.com/blackhorseya/golang-101/pkg/otelx"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const name = "rolldice"
const otelTarget = "localhost:4317"

var tracer trace.Tracer
var meter metric.Meter
var rollCounter metric.Int64Counter

func main() {
	// Initialize OpenTelemetry
	err := otelx.SetupOTelSDK(context.Background(), otelTarget, name)
	if err != nil {
		log.Printf("Failed to initialize OpenTelemetry: %v", err)
		return
	}
	defer func() {
		err = otelx.Shutdown(context.Background())
		if err != nil {
			log.Printf("Failed to shutdown OpenTelemetry: %v", err)
		}
	}()

	// Create a tracer and a meter
	tracer = otel.Tracer(name)
	meter = otel.Meter(name)
	rollCounter, err = meter.Int64Counter(
		"dice.rolls",
		metric.WithDescription("The number of dice rolls"),
		metric.WithUnit("{roll}"),
	)
	if err != nil {
		log.Printf("Failed to create the counter: %v", err)
		return
	}

	// Create a Gin router
	router := gin.Default()
	router.Use(otelgin.Middleware(name))
	router.GET("/rolldice", rolldice)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
		return
	}
}

func rolldice(c *gin.Context) {
	// Start a span
	ctx, span := tracer.Start(c.Request.Context(), "roll")
	defer span.End()

	// Define the number of sides on the die
	sides := 6

	// Generate a random number in the range [0, sides)
	n, err := rand.Int(rand.Reader, big.NewInt(int64(sides)))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate random number"})
		return
	}

	// Add 1 to the result to get a number in the range [1, sides]
	roll := n.Int64() + 1

	rollValueAttr := attribute.Int("roll.value", int(roll))
	span.SetAttributes(rollValueAttr)
	rollCounter.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

	c.JSON(200, gin.H{"roll": roll})
}
