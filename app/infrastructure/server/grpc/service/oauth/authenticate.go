package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
)

func (s *ServiceAuthentication) Authenticate(context context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	return nil, nil
}
