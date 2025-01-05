package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"gorm.io/gorm"
)

type UserDeviceRepository struct {
	db *gorm.DB
}

func NewUserDeviceRepository(db *gorm.DB) *UserDeviceRepository {
	return &UserDeviceRepository{
		db: db,
	}
}

func (r *UserDeviceRepository) WithTx(tx *gorm.DB) *UserDeviceRepository {
	return &UserDeviceRepository{
		db: tx,
	}
}

func (r *UserDeviceRepository) WithContext(ctx context.Context) *UserDeviceRepository {
	return &UserDeviceRepository{
		db: r.db.WithContext(ctx),
	}
}

func (r *UserDeviceRepository) Save(userDevice *entities.UserDevice) error {
	return r.db.Save(userDevice).Error
}

func (r *UserDeviceRepository) FindByUserId(userId uuid.UUID) (*[]entities.UserDevice, error) {
	var userDevices []entities.UserDevice

	if err := r.db.Where("user_id = ?", userId).Find(&userDevices).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &userDevices, nil
}
