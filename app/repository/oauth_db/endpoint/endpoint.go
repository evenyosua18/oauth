package endpoint

import "gorm.io/gorm"

type RepositoryEndpoint struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *RepositoryEndpoint {
	return &RepositoryEndpoint{db: db}
}
