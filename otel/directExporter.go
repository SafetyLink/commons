package otel

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"os"
)

var (
	endpoint      string
	lsAccessToken string
)

func newDirectExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	endpoint = os.Getenv("LIGHTSTEP_URL")
	lsAccessToken = os.Getenv("LIGHTSTEP_ACCESS_TOKEN")
	var headers = map[string]string{
		"lightstep-access-token": lsAccessToken,
	}

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithHeaders(headers),
		otlptracegrpc.WithEndpoint(endpoint),
	)

	return otlptrace.New(ctx, client)
}
