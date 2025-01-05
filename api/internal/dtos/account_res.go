package dtos

import "github.com/google/uuid"

type AccountRes struct {
	ID             uuid.UUID `json:"id"`
	ProfilePicture *string   `json:"profilePicture"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
}
