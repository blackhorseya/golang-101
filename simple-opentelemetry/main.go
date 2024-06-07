package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serviceName = semconv.ServiceNameKey.String("simple-opentelemetry")

func initConn() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:4317", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	return conn, nil
}

func initJaeger(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
) (func(context.Context) error, error) {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create the Jaeger exporter: %w", err)
	}

	processor := trace.NewBatchSpanProcessor(exporter)
	provider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(processor),
	)
	otel.SetTracerProvider(provider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return provider.Shutdown, nil
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
	log.Printf("Waiting for connection...")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := initConn()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	res, err := resource.New(ctx, resource.WithAttributes(serviceName))
	if err != nil {
		log.Println(err)
		return
	}

	jaeger, err := initJaeger(ctx, res, conn)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err = jaeger(ctx); err != nil {
			log.Printf("failed to shutdown Jaeger exporter: %v", err)
			return
		}
	}()

	tracer := otel.Tracer("test-tracer")

	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	fmt.Println("Hello, OpenTelemetry!")
	doWork(ctx)
}

func doWork(ctx context.Context) {
	tracer := otel.Tracer("test-tracer")
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
