package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"time"
)

type Scope struct {
	Id   string `gorm:"size:36;primaryKey"`
	Name string `gorm:"size:50"`

	Endpoints []*Endpoint `gorm:"many2many:scope_endpoints;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (Scope) TableName() string {
	return string(constant.ScopeTable)
}
