package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/repository/oauth_db"
)

type InputPortEndpoint interface {
	GetEndpoints(ctx context.Context, in interface{}) (interface{}, error)
	InsertEndpoint(ctx context.Context, in interface{}) (interface{}, error)
	UpdateEndpoint(ctx context.Context, in interface{}) error
	DeleteEndpoint(ctx context.Context, in interface{}) error
}

type OutputPortEndpoint interface {
	GetEndpointsResponse(in interface{}) (interface{}, error)
	InsertEndpointResponse(in interface{}) (interface{}, error)
}

type InteractionEndpoint struct {
	repo oauth_db.EndpointRepository
	out  OutputPortEndpoint
}

func NewInteractionEndpoint(r oauth_db.EndpointRepository, o OutputPortEndpoint) *InteractionEndpoint {
	return &InteractionEndpoint{
		repo: r,
		out:  o,
	}
}
