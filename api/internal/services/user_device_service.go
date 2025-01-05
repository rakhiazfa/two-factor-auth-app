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

type UserDeviceService struct {
	db                   *gorm.DB
	validator            *utils.Validator
	userDeviceRepository *repositories.UserDeviceRepository
}

func NewUserDeviceService(
	db *gorm.DB,
	validator *utils.Validator,
	userDeviceRepository *repositories.UserDeviceRepository,
) *UserDeviceService {
	return &UserDeviceService{
		db:                   db,
		validator:            validator,
		userDeviceRepository: userDeviceRepository,
	}
}

func (s *UserDeviceService) Create(ctx context.Context, req *dtos.CreateUserDeviceReq) (*entities.UserDevice, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	var userDevice entities.UserDevice

	if err := copier.Copy(&userDevice, req); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.userDeviceRepository.WithTx(tx).Save(&userDevice)
	})
	if err != nil {
		return nil, err
	}

	return &userDevice, nil
}
