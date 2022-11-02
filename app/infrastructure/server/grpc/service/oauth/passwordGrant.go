package oauth

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
	"github.com/evenyosua18/oauth/util/tracer"
)

func (s *ServiceAccessToken) PasswordGrant(context context.Context, in *pb.PasswordGrantRequest) (*pb.AccessTokenResponse, error) {
	//tracer
	ctx, sp := tracer.RootTracer(constant.ServiceLayer, context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call interaction
	res, err := s.uc.PasswordGrant(ctx, in)
	if err != nil {
		tracer.LogError(sp, tracer.CallInteraction, err)
		return nil, err
	}

	tracer.LogResponse(sp, res)
	return res.(*pb.AccessTokenResponse), nil
}
