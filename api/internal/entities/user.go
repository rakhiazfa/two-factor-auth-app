package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseEntityWithSoftDelete
	ProfilePicture *string      `gorm:"type:varchar(100);default:null"`
	Name           string       `gorm:"type:varchar(100)"`
	Username       string       `gorm:"type:varchar(100);unique"`
	Email          string       `gorm:"type:varchar(100);unique"`
	Password       string       `gorm:"type:varchar(100)"`
	UserDevices    []UserDevice `gorm:"foreignKey:UserId;references:ID"`
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		hash, err := m.HashPassword(m.Password)
		if err != nil {
			return err
		}

		m.Password = hash
	}

	return
}

func (m *User) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
