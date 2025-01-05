package entities

import (
	"time"

	"github.com/google/uuid"
)

type TwoFactorAuthSession struct {
	BaseEntity
	UserId                     uuid.UUID
	UserDeviceId               uuid.UUID
	ApprovedBy                 *uuid.UUID
	User                       *User                       `gorm:"foreignKey:UserId;references:ID"`
	UserDevice                 *UserDevice                 `gorm:"foreignKey:UserDeviceId;references:ID"`
	Approver                   *UserDevice                 `gorm:"foreignKey:ApprovedBy;references:ID"`
	CorrectNumber              *string                     `gorm:"type:varchar(100);default:null"`
	Verified                   bool                        `gorm:"type:bool;default:false"`
	ExpiresAt                  time.Time                   `gorm:"type:timestamp"`
	TwoFactorAuthNumberOptions []TwoFactorAuthNumberOption `gorm:"foreignKey:TwoFactorAuthSessionId;references:ID"`
}
