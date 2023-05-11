package otel

import (
	"context"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	tracer := InitTracer()
	ctx := context.Background()

	ctx, span := tracer.Start(ctx, "Testing")
	defer span.End()

	doSomeWork(ctx)

}

func doSomeWork(ctx context.Context) {
	time.Sleep(1 * time.Second)
}
