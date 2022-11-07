package oauth_db

import "context"

type UserRepository interface {
	InsertUser(ctx context.Context, in interface{}) (interface{}, error)
	GetUser(ctx context.Context, in interface{}) (interface{}, error)
}
