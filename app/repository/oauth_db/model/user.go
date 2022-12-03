package model

import (
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          string `gorm:"size:36;primaryKey"`
	Name        string `gorm:"size:100"`
	Password    string `gorm:"size:255"`
	Email       string `gorm:"size:100"`
	Phone       string `gorm:"size:25"`
	Scopes      string `gorm:"size:500"`
	IsSuperRole bool
	IsActive    bool

	RoleId string `gorm:"size:36"`
	Role   Role

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (User) TableName() string {
	return string(constant.UserTable)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	u.Id = uid.String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u User) GetIdColumn() string {
	return "id"
}

func (u User) GetNameColumn() string {
	return "name"
}

func (u User) GetEmailColumn() string {
	return "email"
}

func (u User) GetPhoneColumn() string {
	return "phone"
}
