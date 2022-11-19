package oauth_db

import "context"

type AccessTokenRepository interface {
	InsertAccessToken(ctx context.Context, in interface{}) error
}
