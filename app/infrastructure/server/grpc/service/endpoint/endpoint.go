package endpoint

import "github.com/evenyosua18/oauth/app/usecase/endpoint"

type ServiceEndpoint struct {
	ucEndpoint endpoint.InputPortEndpoint
}

func NewServiceEndpoint(uc endpoint.InputPortEndpoint) *ServiceEndpoint {
	return &ServiceEndpoint{ucEndpoint: uc}
}
