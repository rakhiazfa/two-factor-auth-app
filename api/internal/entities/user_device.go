package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserDevice struct {
	BaseEntityWithSoftDelete
	UserId    uuid.UUID
	User      *User     `gorm:"foreignKey:UserId;references:ID"`
	Type      string    `gorm:"type:varchar(100)"`
	Name      string    `gorm:"type:varchar(100)"`
	Token     string    `gorm:"type:varchar(255)"`
	ExpiresAt time.Time `gorm:"type:timestamp"`
}
