package tracing

import (
	"context"
	"log"
	"pos-graduacao/desafios/observabilidade/input/configs"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitializeTracer(config *configs.Config) func() {
	exporter, err := zipkin.New(config.TRACER_URL)
	if err != nil {
		log.Fatalf("failed to create zipkin exporter: %v", err)
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(config.SERVICE_NAME),
			),
		),
	)
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}
}
