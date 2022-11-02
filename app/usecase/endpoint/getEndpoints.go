package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

func (i *InteractionEndpoint) GetEndpoints(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//get endpoint repo
	endpointResponse, err := i.repo.GetEndpoints(ctx, in)
	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}

	//mapping response
	var endpoints []entity.Endpoint
	if err := mapstructure.Decode(endpointResponse, &endpoints); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	tracer.LogResponse(sp, endpoints)
	return i.out.GetEndpointsResponse(&entity.GetEndpointsResponse{Endpoints: endpoints})
}
