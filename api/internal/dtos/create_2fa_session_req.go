package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Create2FASessionReq struct {
	UserId       uuid.UUID `json:"userId" validate:"required,uuid"`
	UserDeviceId uuid.UUID `json:"userDeviceId" validate:"required,uuid"`
	Verified     bool      `json:"verified" validate:"boolean"`
	ExpiresAt    time.Time `json:"expiresAt" validate:"required"`
}
