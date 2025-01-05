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

type AccountService struct {
	db                   *gorm.DB
	validator            *utils.Validator
	userDeviceRepository *repositories.UserDeviceRepository
}

func NewAccountService(
	db *gorm.DB,
	validator *utils.Validator,
	userDeviceRepository *repositories.UserDeviceRepository,
) *AccountService {
	return &AccountService{
		db:                   db,
		validator:            validator,
		userDeviceRepository: userDeviceRepository,
	}
}

func (s *AccountService) FindAccountDevices(ctx context.Context, authUser *entities.User) (*dtos.ListUserDeviceRes, error) {
	userDevices, err := s.userDeviceRepository.FindByUserId(authUser.ID)
	if err != nil {
		return nil, err
	}

	var userDevicesRes dtos.ListUserDeviceRes

	if err := copier.Copy(&userDevicesRes, userDevices); err != nil {
		return nil, err
	}

	return &userDevicesRes, nil
}
