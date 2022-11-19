package user

import "gorm.io/gorm"

type RepositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db: db}
}
