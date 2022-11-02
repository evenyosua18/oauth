package endpoint

import "github.com/jinzhu/gorm"

type RepositoryEndpoint struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *RepositoryEndpoint {
	return &RepositoryEndpoint{db: db}
}
