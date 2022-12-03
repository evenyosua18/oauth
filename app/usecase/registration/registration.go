package registration

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/repository/oauth_db"
)

type InputPortRegistration interface {
	RegistrationUser(ctx context.Context, in interface{}) (interface{}, error)
}

type OutputPortRegistration interface {
	RegistrationUserResponse(in interface{}) (interface{}, error)
}

type InteractionRegistration struct {
	DefaultRoleId string

	repo oauth_db.UserRepository
	out  OutputPortRegistration
}

func NewInteractionRegistration(r oauth_db.UserRepository, o OutputPortRegistration) *InteractionRegistration {
	//set default role id
	defaultRole := "d5e54e04-6def-43b4-ac99-800d315665c4" //next, it should be from redis

	return &InteractionRegistration{
		repo:          r,
		out:           o,
		DefaultRoleId: defaultRole,
	}
}
