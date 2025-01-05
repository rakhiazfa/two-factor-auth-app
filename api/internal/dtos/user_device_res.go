package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserDeviceRes struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ListUserDeviceRes = []UserDeviceRes
