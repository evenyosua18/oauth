package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/util/tracer"
)

type DeleteEndpointRequest struct {
	Id        string
	DeletedAt string
}

func (i *InteractionEndpoint) DeleteEndpoint(context context.Context, in interface{}) error {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call repository
	if err := i.repo.DeleteEndpoint(ctx, in); err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return err
	}

	return nil
}
