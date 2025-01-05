package dtos

import "github.com/google/uuid"

type CreateUserDeviceReq struct {
	UserId uuid.UUID `json:"userId" validate:"required,uuid"`
	Type   string    `json:"type" validate:"required,max=100"`
	Name   string    `json:"name" validate:"required,max=100"`
	Token  string    `json:"token" validate:"required,max=255"`
}
