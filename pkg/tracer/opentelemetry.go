package tracer

import (
	"fmt"
	"github.com/prok05/spot-instrument-service/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

import "context"

func New(ctx context.Context, cfg *config.Config) (func(context.Context) error, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(semconv.ServiceName(cfg.App.Name)))
	if err != nil {
		return nil, fmt.Errorf("tracer - New - resource.New: %w", err)
	}

	conn, err := grpc.NewClient(cfg.Jaeger.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("tracer - New - grpc.NewClient: %w", err)
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("tracer - New - otlptracegrpc.New: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return provider.Shutdown, nil
}
