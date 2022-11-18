package accessToken

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/repository/oauth_db"
)

type InputPortAccessToken interface {
	PasswordGrant(ctx context.Context, in interface{}) (interface{}, error)
}

type OutputPortAccessToken interface {
	AccessTokenResponse(in interface{}) (interface{}, error)
}

type InteractionAccessToken struct {
	repo oauth_db.OauthClientRepository
	user oauth_db.UserRepository
	out  OutputPortAccessToken
}

func NewInteractionAccessToken(r oauth_db.OauthClientRepository, u oauth_db.UserRepository, o OutputPortAccessToken) *InteractionAccessToken {
	return &InteractionAccessToken{
		repo: r,
		user: u,
		out:  o,
	}
}
