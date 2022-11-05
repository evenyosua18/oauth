package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *ServiceEndpoint) UpdateEndpoint(context context.Context, in *pb.UpdateEndpointRequest) (*empty.Empty, error) {
	//tracer
	ctx, sp := tracer.RootTracer(constant.ServiceLayer, context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call interaction
	if err := s.ucEndpoint.UpdateEndpoint(ctx, in); err != nil {
		tracer.LogError(sp, tracer.CallInteraction, err)
		return nil, err
	}

	return &empty.Empty{}, nil
}
