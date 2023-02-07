package otel

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	resource, _ := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("s3-poc"),
		))

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	connection, _ := grpc.DialContext(ctx, "localhost:4317", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	tracerExporter, _ := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(connection))

	batachSpan := sdktrace.NewBatchSpanProcessor(tracerExporter)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource),
		sdktrace.WithSpanProcessor(batachSpan),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tracerProvider.Shutdown, nil
}
