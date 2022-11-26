package authentication

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/repository/oauth_db"
)

type InputPortAuthentication interface {
	Authenticate(ctx context.Context, in interface{}) error
}

type InteractionAuthentication struct {
	user        oauth_db.UserRepository
	accessToken oauth_db.AccessTokenRepository
}

func NewInteractionAuthentication(u oauth_db.UserRepository, at oauth_db.AccessTokenRepository) *InteractionAuthentication {
	return &InteractionAuthentication{
		user:        u,
		accessToken: at,
	}
}
