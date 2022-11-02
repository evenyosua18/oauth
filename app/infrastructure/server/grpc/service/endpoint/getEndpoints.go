package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
	"github.com/evenyosua18/oauth/util/tracer"
)

func (s *ServiceEndpoint) GetEndpoints(context context.Context, in *pb.GetEndpointsRequest) (*pb.GetEndpointsResponse, error) {
	//tracer
	ctx, sp := tracer.RootTracer(constant.ServiceLayer, context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call interaction
	res, err := s.ucEndpoint.GetEndpoints(ctx, in)
	if err != nil {
		tracer.LogError(sp, tracer.CallInteraction, err)
		return nil, err
	}

	tracer.LogResponse(sp, res)
	return res.(*pb.GetEndpointsResponse), nil
}
