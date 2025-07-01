package interceptor

import (
	"context"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Interceptor) Panic() pbgrpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *pbgrpc.UnaryServerInfo,
		handler pbgrpc.UnaryHandler) (resp any, err error) {

		defer func() {
			if r := recover(); r != nil {
				i.L.Error(
					"grpc interceptor", info.FullMethod,
					"request failed", "panic",
					"error", r,
				)
				err = status.Errorf(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}
