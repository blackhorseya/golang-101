package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initJaeger() func() {
	// Create the Jaeger exporter
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		panic(err)
	}

	// Create the Jaeger trace provider
	tracer := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("simple-opentelemetry"),
		)),
	)
	otel.SetTracerProvider(tracer)

	return func() {
		_ = tracer.Shutdown(context.Background())
	}
}

func initTracer() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	tracer := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("simple-opentelemetry"),
		)),
	)
	otel.SetTracerProvider(tracer)

	return func() {
		_ = tracer.Shutdown(context.Background())
	}
}

func main() {
	cleanup := initJaeger()
	defer cleanup()

	tracer := otel.Tracer("simple-opentelemetry")

	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	fmt.Println("Hello, OpenTelemetry!")
	doWork(ctx)
}

func doWork(ctx context.Context) {
	tracer := otel.Tracer("simple-opentelemetry")
	_, span := tracer.Start(ctx, "doWork")
	defer span.End()

	// random sleep to simulate work
	fmt.Println("Working...")
	duration := getRandomDuration(1, 5)
	time.Sleep(duration)
}

func getRandomDuration(min, max int) time.Duration {
	// Create a buffer to store the random bytes
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// Convert the random bytes to an integer
	randomInt := binary.BigEndian.Uint64(b)

	// Scale the integer to the desired range and convert to a time.Duration
	randomDuration := min + int(randomInt%(uint64(max-min+1)))
	return time.Duration(randomDuration) * time.Second
}
