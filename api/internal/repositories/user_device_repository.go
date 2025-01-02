package repositories

import (
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

func (r *UserDeviceRepository) Save(user *entities.UserDevice) error {
	return r.db.Save(user).Error
}
