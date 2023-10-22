package middleware

import (
	"context"
	"errors"
	"github.com/pramudyas69/Go-Grpc/server/internal/config"
	"github.com/pramudyas69/Go-Grpc/server/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Authenticate(cnf *config.Config, tokenKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if info.FullMethod == "/proto.UserService/Login" || info.FullMethod == "/proto.UserService/Register" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("metadata not found")
		}

		token := md.Get(tokenKey)
		if len(token) == 0 {
			return nil, errors.New("token not found")
		}

		claims, err := util.ValidateToken(cnf, token[0])
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "x-user", claims.UserID)

		return handler(ctx, req)
	}
}
