package services

import (
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (service *AuthService) SignIn(req dtos.SignInReq) (string, error) {
	return "", nil
}

func (service *AuthService) SignUp(req dtos.SignUpReq) error {
	return nil
}
