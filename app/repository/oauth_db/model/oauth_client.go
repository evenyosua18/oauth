package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"time"
)

type OauthClient struct {
	Id           string `gorm:"size:36;primaryKey"`
	ClientId     string `gorm:"size:8"`
	URI          string `gorm:"size:100"`
	ClientSecret string `gorm:"size:16"`
	Scopes       string `gorm:"size:150"`
	ClientType   string `gorm:"size:25"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (OauthClient) TableName() string {
	return string(constant.OauthClientTable)
}

func (*OauthClient) GetClientIdColumn() string {
	return "client_id"
}
