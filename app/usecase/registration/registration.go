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
	repo oauth_db.UserRepository
	out  OutputPortRegistration
}

func NewInteractionRegistration(r oauth_db.UserRepository, o OutputPortRegistration) *InteractionRegistration {
	return &InteractionRegistration{
		repo: r,
		out:  o,
	}
}
