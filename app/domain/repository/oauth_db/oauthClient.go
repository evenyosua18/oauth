package oauth_db

import "context"

type OauthClientRepository interface {
	GetOauthClient(ctx context.Context, in interface{}) (interface{}, error)
}
