package oauth_client

import "gorm.io/gorm"

type RepositoryOauthClient struct {
	db *gorm.DB
}

func NewOauthClientRepository(db *gorm.DB) *RepositoryOauthClient {
	return &RepositoryOauthClient{db: db}
}
