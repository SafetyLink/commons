package otel

import (
	"context"
	"github.com/joho/godotenv"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	//load env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	tracer := InitTracer()
	ctx := context.Background()

	ctx, span := tracer.Start(ctx, "Testing")
	defer span.End()

	doSomeWork(ctx)

}

func doSomeWork(ctx context.Context) {
	time.Sleep(1 * time.Second)
}
