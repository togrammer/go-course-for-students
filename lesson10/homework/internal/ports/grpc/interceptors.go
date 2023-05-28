package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func LoggerInterceptor(ctx context.Context, r interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, r)
	log.Println("method: ", info.FullMethod, "error: ", err)
	return res, err
}

func PanicInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic occurred:", r)
			err = status.Error(codes.Internal, "Internal server error")
		}
	}()
	return handler(ctx, req)
}
