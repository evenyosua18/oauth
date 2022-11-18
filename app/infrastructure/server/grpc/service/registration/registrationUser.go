package registration

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/evenyosua18/oauth/util/tracer"
)

func (s *ServiceRegistration) RegistrationUser(context context.Context, in *pb.RegistrationUserRequest) (*pb.RegistrationUserResponse, error) {
	//tracer
	ctx, sp := tracer.RootTracer(constant.ServiceLayer, context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call interaction
	res, err := s.uc.RegistrationUser(ctx, in)
	if err != nil {
		tracer.LogError(sp, tracer.CallInteraction, err)
		return nil, err
	}

	tracer.LogResponse(sp, res)
	return res.(*pb.RegistrationUserResponse), nil
}