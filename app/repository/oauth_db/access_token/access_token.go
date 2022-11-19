package access_token

import "gorm.io/gorm"

type RepositoryAccessToken struct {
	db *gorm.DB
}

func NewAccessTokenRepository(db *gorm.DB) *RepositoryAccessToken {
	return &RepositoryAccessToken{db: db}
}
