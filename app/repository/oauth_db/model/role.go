package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"reflect"
	"time"
)

type Role struct {
	Id          string `gorm:"size:36;primaryKey"`
	RoleName    string `gorm:"size:50"`
	Scopes      string `gorm:"size:150"`
	Description string `gorm:"size:100"`
	IsSuperRole bool

	Endpoints []*Endpoint `gorm:"many2many:role_endpoints;"`
	Users     []User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (Role) TableName() string {
	return string(constant.RoleTable)
}

func (r Role) JoinName() string {
	return reflect.TypeOf(r).Name()
}
