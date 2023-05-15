package otel

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func RecordError(customError error, err error, message string, span trace.Span) error {
	span.RecordError(err)
	span.SetStatus(codes.Error, message)

	return customError
}

func RecordErrorWithAttribute(customError error, err error, message string, span trace.Span, kv ...attribute.KeyValue) error {
	span.RecordError(err)
	span.SetStatus(codes.Error, message)
	span.SetAttributes(kv...)

	return customError
}
