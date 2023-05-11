package otel

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func RecordError(customError error, err error, message string, span trace.Span, logger *zap.Logger) error {
	span.RecordError(err)
	span.SetStatus(codes.Error, message)
	logger.Error(message, zap.Error(err))

	return customError
}

func RecordErrorWithAttribute(customError error, err error, message string, span trace.Span, logger *zap.Logger, kv ...attribute.KeyValue) error {
	span.RecordError(err)
	span.SetStatus(codes.Error, message)
	span.SetAttributes(kv...)
	logger.Error(message, zap.Error(err))

	return customError
}
