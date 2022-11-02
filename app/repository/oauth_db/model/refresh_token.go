package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"time"
)

type RefreshToken struct {
	Id           string `gorm:"size:36;primaryKey"`
	RefreshToken string `gorm:"size:36"`

	AccessTokenId string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (RefreshToken) TableName() string {
	return string(constant.RefreshTokenTable)
}
