package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Endpoint struct {
	Id           string `gorm:"size:36;primaryKey"`
	Name         string `gorm:"size:50"`
	EndpointType string `gorm:"size:4"`
	Description  string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (Endpoint) TableName() string {
	return string(constant.EndpointTable)
}

func (e *Endpoint) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	e.Id = uid.String()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Endpoint) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return
}

func (Endpoint) GetIdColumn() string {
	return "id"
}

func (Endpoint) GetEndpointTypeColumn() string {
	return "endpoint_type"
}

func (Endpoint) GetNameColumn() string {
	return "name"
}
