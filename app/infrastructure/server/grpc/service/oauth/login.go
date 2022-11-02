package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
)

func (s *ServiceAccessToken) Login(context context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
