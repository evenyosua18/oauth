package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
)

func (s *ServiceAccessToken) AuthorizationCode(context context.Context, request *pb.AuthorizationCodeRequest) (*pb.AccessTokenResponse, error) {
	return nil, nil
}
