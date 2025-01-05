package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"gorm.io/gorm"
)

type TwoFactorAuthNumberOptionRepository struct {
	db *gorm.DB
}

func NewTwoFactorAuthNumberOptionRepository(db *gorm.DB) *TwoFactorAuthNumberOptionRepository {
	return &TwoFactorAuthNumberOptionRepository{
		db: db,
	}
}

func (r *TwoFactorAuthNumberOptionRepository) WithTx(tx *gorm.DB) *TwoFactorAuthNumberOptionRepository {
	return &TwoFactorAuthNumberOptionRepository{
		db: tx,
	}
}

func (r *TwoFactorAuthNumberOptionRepository) WithContext(ctx context.Context) *TwoFactorAuthNumberOptionRepository {
	return &TwoFactorAuthNumberOptionRepository{
		db: r.db.WithContext(ctx),
	}
}

func (r *TwoFactorAuthNumberOptionRepository) CreateMany(twoFactorAuthNumberOptions *[]entities.TwoFactorAuthNumberOption) error {
	return r.db.Create(twoFactorAuthNumberOptions).Error
}

func (r *TwoFactorAuthNumberOptionRepository) FindOneById(id uuid.UUID) (*entities.TwoFactorAuthNumberOption, error) {
	var twoFactorAuthNumberOption entities.TwoFactorAuthNumberOption

	if err := r.db.First(&twoFactorAuthNumberOption, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &twoFactorAuthNumberOption, nil
}
