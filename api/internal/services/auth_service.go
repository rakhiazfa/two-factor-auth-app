package services

import (
	"context"
	"net/http"
	"time"

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
	db                   *gorm.DB
	validator            *utils.Validator
	userRepository       *repositories.UserRepository
	userDeviceService    *UserDeviceService
	twoFactorAuthService *TwoFactorAuthService
}

func NewAuthService(
	db *gorm.DB,
	validator *utils.Validator,
	userRepository *repositories.UserRepository,
	userDeviceService *UserDeviceService,
	twoFactorAuthService *TwoFactorAuthService,
) *AuthService {
	return &AuthService{
		db:                   db,
		validator:            validator,
		userRepository:       userRepository,
		userDeviceService:    userDeviceService,
		twoFactorAuthService: twoFactorAuthService,
	}
}

func (s *AuthService) SignIn(ctx context.Context, req *dtos.SignInReq) (uuid.UUID, string, error) {
	if err := s.validator.Validate(req); err != nil {
		return uuid.Nil, "", err
	}

	user, err := s.ValidateUserCredentials(ctx, req.UsernameOrEmail, req.Password)
	if err != nil {
		return uuid.Nil, "", err
	}

	userDevice, err := s.userDeviceService.Create(ctx, &dtos.CreateUserDeviceReq{
		UserId: user.ID,
		Type:   req.DeviceType,
		Name:   req.DeviceName,
		Token:  req.DeviceToken,
	})
	if err != nil {
		return uuid.Nil, "", err
	}

	twoFactorAuthSession, err := s.twoFactorAuthService.Create2FASession(ctx, &dtos.Create2FASessionReq{
		UserId:       user.ID,
		UserDeviceId: userDevice.ID,
		Verified:     false,
		ExpiresAt:    time.Now().Add(1 * time.Minute),
	})
	if err != nil {
		return uuid.Nil, "", err
	}

	correctNumber, err := s.twoFactorAuthService.Create2FANumberOptions(ctx, twoFactorAuthSession)
	if err != nil {
		return uuid.Nil, "", err
	}

	return twoFactorAuthSession.ID, correctNumber, nil
}

func (s *AuthService) SignUp(ctx context.Context, req *dtos.SignUpReq) (string, string, error) {
	if err := s.validator.Validate(req); err != nil {
		return "", "", err
	}

	err := s.validateUserEmailAndUsername(ctx, req)
	if err != nil {
		return "", "", err
	}

	var user entities.User

	if err := copier.Copy(&user, req); err != nil {
		return "", "", err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.userRepository.WithTx(tx).Save(&user)
	})
	if err != nil {
		return "", "", err
	}

	userDevice, err := s.userDeviceService.Create(ctx, &dtos.CreateUserDeviceReq{
		UserId: user.ID,
		Type:   req.DeviceType,
		Name:   req.DeviceName,
		Token:  req.DeviceToken,
	})
	if err != nil {
		return "", "", err
	}

	twoFactorAuthSession, err := s.twoFactorAuthService.Create2FASession(ctx, &dtos.Create2FASessionReq{
		UserId:       user.ID,
		UserDeviceId: userDevice.ID,
		Verified:     true,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(jwt.MapClaims{
		"sub": twoFactorAuthSession.ID,
	})
	if err != nil {
		return "", "", err
	}

	accessToken, err := utils.CreateAccessToken(jwt.MapClaims{
		"sub": twoFactorAuthSession.ID,
	})
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (s *AuthService) ValidateUserCredentials(ctx context.Context, usernameOrEmail string, password string) (*entities.User, error) {
	user, err := s.userRepository.WithContext(ctx).FindOneByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, utils.NewHttpError(http.StatusUnauthorized, "Unauthorized", nil)
	}

	return user, nil
}

func (s *AuthService) validateUserEmailAndUsername(ctx context.Context, req *dtos.SignUpReq) error {
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

	return nil
}
