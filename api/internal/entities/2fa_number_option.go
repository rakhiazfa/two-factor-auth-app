package entities

import "github.com/google/uuid"

type TwoFactorAuthNumberOption struct {
	BaseEntity
	TwoFactorAuthSessionId uuid.UUID
	TwoFactorAuthSession   *TwoFactorAuthSession `gorm:"foreignKey:TwoFactorAuthSessionId;references:ID"`
	Number                 string                `gorm:"type:varchar(100)"`
}
