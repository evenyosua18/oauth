package accessToken

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/repository/oauth_db"
	"github.com/evenyosua18/oauth/config"
)

const (
	defaultExpiredTime = "1"
)

type InputPortAccessToken interface {
	PasswordGrant(ctx context.Context, in interface{}) (interface{}, error)
}

type OutputPortAccessToken interface {
	AccessTokenResponse(in interface{}) (interface{}, error)
}

type InteractionAccessToken struct {
	ExpiredTime string

	repo     oauth_db.OauthClientRepository
	user     oauth_db.UserRepository
	accToken oauth_db.AccessTokenRepository
	out      OutputPortAccessToken
}

func NewInteractionAccessToken(r oauth_db.OauthClientRepository, u oauth_db.UserRepository, at oauth_db.AccessTokenRepository, o OutputPortAccessToken) *InteractionAccessToken {
	interaction := InteractionAccessToken{
		repo:     r,
		user:     u,
		out:      o,
		accToken: at,
	}

	//expired time
	interaction.ExpiredTime = config.GetConfig().Server.Token.Expired

	if interaction.ExpiredTime == "" {
		interaction.ExpiredTime = defaultExpiredTime
	}

	return &interaction
}
