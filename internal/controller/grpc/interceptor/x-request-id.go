package interceptor

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (i *Interceptor) XRequestID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		var requestID string
		if ok {
			ids := md.Get(string(ContextRequestID))
			if len(ids) > 0 {
				requestID = ids[0]
			}
		}
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx = context.WithValue(ctx, ContextRequestID, requestID)

		return handler(ctx, req)
	}
}
