package oauth_client

import "github.com/jinzhu/gorm"

type RepositoryOauthClient struct {
	db *gorm.DB
}

func NewOauthClientRepository(db *gorm.DB) *RepositoryOauthClient {
	return &RepositoryOauthClient{db: db}
}
