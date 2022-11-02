package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *ServiceAuthentication) Logout(context context.Context, request *pb.LogoutRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
