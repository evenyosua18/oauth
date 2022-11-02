package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/util/tracer"
)

func (i *InteractionEndpoint) UpdateEndpoint(context context.Context, in interface{}) error {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//call repository
	if err := i.repo.UpdateEndpoint(ctx, in); err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return err
	}

	return nil
}
