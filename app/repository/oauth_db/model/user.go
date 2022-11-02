package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"time"
)

type User struct {
	Id       string `gorm:"size:36;primaryKey"`
	Name     string `gorm:"size:100"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:100"`
	Phone    string `gorm:"size:25"`
	IsActive bool

	RoleId string `gorm:"size:36"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (User) TableName() string {
	return string(constant.UserTable)
}
