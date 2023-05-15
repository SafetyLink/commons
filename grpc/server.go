package grpc

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			otelgrpc.StreamServerInterceptor(),
		))

	return server
}
