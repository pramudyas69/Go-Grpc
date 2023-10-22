package rpc

import (
	"context"
	"github.com/pramudyas69/Go-Grpc/server/domain"
	"github.com/pramudyas69/Go-Grpc/server/dto"
	pb "github.com/pramudyas69/Go-Grpc/server/proto"
	"google.golang.org/grpc"
)

type userGrpc struct {
	pb.UnimplementedUserServiceServer
	userSvc domain.UserService
}

func NewUser(grpcServer *grpc.Server, userSvc domain.UserService) {
	userRpc := &userGrpc{
		userSvc: userSvc,
	}
	pb.RegisterUserServiceServer(grpcServer, userRpc)
}

func (u userGrpc) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.Empty, error) {
	user := dto.RegisterReq{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	_, err := u.userSvc.Register(ctx, user)
	if err != nil {
		return nil, nil
	}

	return &pb.Empty{}, nil
}

func (u userGrpc) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	req := dto.LoginReq{
		Username: request.Username,
		Password: request.Password,
	}

	res, err := u.userSvc.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: res.Token,
	}, nil
}

func (u userGrpc) ValidateToken(ctx context.Context,
	request *pb.ValidateTokenRequest) (*pb.Empty, error) {
	_, err := u.userSvc.ValidateToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
