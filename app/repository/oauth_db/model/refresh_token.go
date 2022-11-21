package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (e *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if e.Id == "" {
		uid, err := uuid.NewUUID()

		if err != nil {
			return err
		}

		e.Id = uid.String()
	}

	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return nil
}
