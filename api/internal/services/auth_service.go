package services

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db             *gorm.DB
	validator      *utils.Validator
	userRepository *repositories.UserRepository
}

func NewAuthService(db *gorm.DB, validator *utils.Validator, userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		db:             db,
		validator:      validator,
		userRepository: userRepository,
	}
}

func (s *AuthService) SignIn(ctx context.Context, req *dtos.SignInReq) (string, error) {
	if err := s.validator.Validate(req); err != nil {
		return "", err
	}

	return "", nil
}

func (s *AuthService) SignUp(ctx context.Context, req *dtos.SignUpReq) error {
	if err := s.validator.Validate(req); err != nil {
		return err
	}

	userWithSameEmail, err := s.userRepository.FindByEmailUnscoped(req.Email)
	if err != nil {
		return err
	}

	if userWithSameEmail != nil {
		return utils.NewUniqueFieldError("email", "An account with this email already exists", nil)
	}

	var user entities.User

	if err := copier.Copy(&user, &req); err != nil {
		return err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.userRepository.WithTx(tx).Save(&user)
	})
	if err != nil {
		return err
	}

	return nil
}
