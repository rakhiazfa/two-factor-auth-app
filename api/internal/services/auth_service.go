package services

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"golang.org/x/crypto/bcrypt"
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

func (s *AuthService) SignIn(ctx context.Context, req *dtos.SignInReq) (string, string, error) {
	if err := s.validator.Validate(req); err != nil {
		return "", "", err
	}

	user, err := s.userRepository.WithContext(ctx).FindOneByUsernameOrEmail(req.UsernameOrEmail)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
	}

	log.Println(user.UserDevices)

	jti := uuid.New()

	refreshToken, err := utils.CreateRefreshToken(jwt.MapClaims{
		"sub": user.ID,
		"jti": jti,
	})
	if err != nil {
		return "", "", err
	}

	accessToken, err := utils.CreateAccessToken(jwt.MapClaims{
		"sub": user.ID,
		"jti": jti,
	})
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (s *AuthService) SignUp(ctx context.Context, req *dtos.SignUpReq) error {
	if err := s.validator.Validate(req); err != nil {
		return err
	}

	userWithSameUsername, err := s.userRepository.WithContext(ctx).FindOneByUsernameUnscoped(req.Username)
	if err != nil {
		return err
	}
	if userWithSameUsername != nil {
		return utils.NewUniqueFieldError("username", "An account with this username already exists", nil)
	}

	userWithSameEmail, err := s.userRepository.WithContext(ctx).FindOneByEmailUnscoped(req.Email)
	if err != nil {
		return err
	}
	if userWithSameEmail != nil {
		return utils.NewUniqueFieldError("email", "An account with this email already exists", nil)
	}

	var user entities.User

	if err := copier.Copy(&user, req); err != nil {
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
