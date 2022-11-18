package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"time"
)

type AccessToken struct {
	Id            string `gorm:"size:36;primaryKey"`
	IpAddress     string `gorm:"size:50"`
	ExpireAt      time.Time
	UserId        string `gorm:"size:36"`
	OauthClientId string `gorm:"size:36"`

	RefreshTokens []RefreshToken `gorm:"foreignKey:AccessTokenId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (AccessToken) TableName() string {
	return string(constant.AccessTokenTable)
}
