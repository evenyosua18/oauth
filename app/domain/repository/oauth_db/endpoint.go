package oauth_db

import "context"

type EndpointRepository interface {
	GetEndpoints(ctx context.Context, in interface{}) (interface{}, error)
	InsertEndpoint(ctx context.Context, in interface{}) (interface{}, error)
	UpdateEndpoint(ctx context.Context, in interface{}) error
	DeleteEndpoint(ctx context.Context, in interface{}) error
}
