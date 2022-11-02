package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/util/tracer"
)

func (i *InteractionEndpoint) InsertEndpoint(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call repo
	res, err := i.repo.InsertEndpoint(ctx, in)
	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}

	tracer.LogResponse(sp, res)
	return i.out.InsertEndpointResponse(res)
}
