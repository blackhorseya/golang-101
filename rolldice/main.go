package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const name = "rolldice"

var serviceName = semconv.ServiceNameKey.String(name)

var tracer trace.Tracer
var meter metric.Meter

func initConn() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:4317", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	return conn, nil
}

func newTracer(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create the Jaeger exporter: %w", err)
	}

	processor := sdktrace.NewBatchSpanProcessor(exporter)
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)
	otel.SetTracerProvider(provider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	tracer = provider.Tracer(name)

	return provider, nil
}

func newMeter(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
) (p *sdkmetric.MeterProvider, err error) {
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create the OTLP exporter: %w", err)
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(3*time.Second))),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(provider)

	meter = provider.Meter(name)

	return provider, nil
}

func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	res, err := resource.New(ctx, resource.WithAttributes(serviceName))
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	conn, err := initConn()
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	traceProvider, err := newTracer(ctx, res, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Jaeger: %w", err)
	}
	shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)

	meterProvider, err := newMeter(ctx, res, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OTLP: %w", err)
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)

	return shutdown, nil
}

func main() {
	ctx := context.Background()

	shutdown, err := setupOTelSDK(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err = shutdown(ctx); err != nil {
			log.Printf("Failed to shutdown the provider: %v", err)
		}
	}()

	router := gin.Default()
	router.Use(otelgin.Middleware(name))
	router.GET("/rolldice", rolldice)

	err = router.Run(":8080")
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
		return
	}
}

func rolldice(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "roll")
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

	span.SetAttributes(attribute.Int("roll.value", int(roll)))

	c.JSON(200, gin.H{"roll": roll})
}
