package user

import "github.com/jinzhu/gorm"

type RepositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db: db}
}
