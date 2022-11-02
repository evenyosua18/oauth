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
	out  OutputPortAccessToken
}

func NewInteractionAccessToken(r oauth_db.OauthClientRepository, o OutputPortAccessToken) *InteractionAccessToken {
	return &InteractionAccessToken{
		repo: r,
		out:  o,
	}
}
