package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/evenyosua18/oauth/util/tracer"
)

func (s *ServiceEndpoint) InsertEndpoint(context context.Context, in *pb.InsertEndpointRequest) (*pb.InsertEndpointResponse, error) {
	//tracer
	ctx, sp := tracer.RootTracer(constant.ServiceLayer, context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call interaction
	res, err := s.ucEndpoint.InsertEndpoint(ctx, in)
	if err != nil {
		tracer.LogError(sp, tracer.CallInteraction, err)
		return nil, err
	}

	tracer.LogResponse(sp, res)
	return res.(*pb.InsertEndpointResponse), nil
}
