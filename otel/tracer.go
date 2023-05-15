package otel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
)

var (
	serviceName    string
	serviceVersion string
	collectorAddr  string
	lsEnvironment  string
)

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	collectorAddr = os.Getenv("COLLECTOR_URL")

	exporter, err :=
		otlptracegrpc.New(ctx,
			// WithInsecure lets us use http instead of https (for local dev only).
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(collectorAddr),
		)

	return exporter, err
}

func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	serviceName = os.Getenv("LIGHTSTEP_SERVICE_NAME")
	serviceVersion = os.Getenv("LIGHTSTEP_SERVICE_VERSION")
	lsEnvironment = os.Getenv("LIGHTSTEP_ENVIRONMENT")

	resource, rErr :=
		resource.Merge(
			resource.Default(),
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				semconv.ServiceVersionKey.String(serviceVersion),
				attribute.String("environment", lsEnvironment),
			),
		)

	if rErr != nil {
		panic(rErr)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)
}

func InitTracer() trace.Tracer {
	ctx := context.Background()

	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp)

	//defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return tp.Tracer(serviceName, trace.WithInstrumentationVersion(serviceVersion))
}
