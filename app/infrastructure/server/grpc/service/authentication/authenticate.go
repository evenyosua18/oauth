package authentication

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *ServiceAuthentication) Authenticate(context context.Context, in *pb.AuthenticateRequest) (*empty.Empty, error) {
	//tracer
	ctx, sp := tracer.RootTracer(constant.ServiceLayer, context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call interaction
	if err := s.uc.Authenticate(ctx, in); err != nil {
		tracer.LogError(sp, tracer.CallInteraction, err)
		return nil, err
	}

	return &empty.Empty{}, nil
}
