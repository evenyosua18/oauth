package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
)

func (s *ServiceAccessToken) ClientCredential(context context.Context, request *pb.ClientCredentialRequest) (*pb.AccessTokenResponse, error) {
	return nil, nil
}
