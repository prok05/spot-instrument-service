package interceptor

import (
	"context"
	pbgrpc "google.golang.org/grpc"
)

func (i *Interceptor) Log() pbgrpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *pbgrpc.UnaryServerInfo,
		handler pbgrpc.UnaryHandler) (resp any, err error) {

		reqID := ctx.Value(ContextRequestID)

		m, err := handler(ctx, req)
		if err != nil {
			i.L.Error("grpc interceptor",
				info.FullMethod, "request failed",
				"request_id", reqID)
		} else {
			i.L.Info("grpc interceptor",
				info.FullMethod, "request success",
				"request_id", reqID)
		}

		return m, err
	}
}
