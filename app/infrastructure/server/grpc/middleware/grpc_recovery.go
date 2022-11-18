package middleware

import (
	"context"
	"github.com/evenyosua18/oauth/util/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				logger.Log(logger.Fatal, info.FullMethod, req)
				err = status.Errorf(codes.Internal, "PANIC | request data : %v", req)
			}
		}()

		resp, err = handler(ctx, req)
		panicked = false
		return resp, err
	}
}
