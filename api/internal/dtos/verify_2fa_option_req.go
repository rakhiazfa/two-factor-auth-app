package dtos

import "github.com/google/uuid"

type Verify2FAOptionReq struct {
	OptionId uuid.UUID `json:"optionId" validate:"required,uuid"`
}
