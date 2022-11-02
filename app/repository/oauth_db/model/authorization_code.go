package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"time"
)

type AuthorizationCode struct {
	Id           string `gorm:"size:36;primaryKey"`
	ExpireAt     time.Time
	IpAddress    string `gorm:"size:50"`
	ClientSecret string `gorm:"size:16"`
	ClientId     string `gorm:"size:8"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (AuthorizationCode) TableName() string {
	return string(constant.AuthorizationCodeTable)
}
