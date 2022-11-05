package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
)

func (s *ServiceAccessToken) CheckClient(context context.Context, request *pb.CheckClientRequest) (*pb.CheckClientResponse, error) {
	return nil, nil
}
