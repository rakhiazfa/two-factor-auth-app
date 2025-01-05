package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"gorm.io/gorm"
)

type TwoFactorAuthSessionRepository struct {
	db *gorm.DB
}

func NewTwoFactorAuthSessionRepository(db *gorm.DB) *TwoFactorAuthSessionRepository {
	return &TwoFactorAuthSessionRepository{
		db: db,
	}
}

func (r *TwoFactorAuthSessionRepository) WithTx(tx *gorm.DB) *TwoFactorAuthSessionRepository {
	return &TwoFactorAuthSessionRepository{
		db: tx,
	}
}

func (r *TwoFactorAuthSessionRepository) WithContext(ctx context.Context) *TwoFactorAuthSessionRepository {
	return &TwoFactorAuthSessionRepository{
		db: r.db.WithContext(ctx),
	}
}

func (r *TwoFactorAuthSessionRepository) Create(twoFactorAuthSession *entities.TwoFactorAuthSession) error {
	return r.db.Create(twoFactorAuthSession).Error
}

func (r *TwoFactorAuthSessionRepository) FindOneById(id uuid.UUID, relations ...string) (*entities.TwoFactorAuthSession, error) {
	var twoFactorAuthSession entities.TwoFactorAuthSession

	q := r.db

	for _, relation := range relations {
		q = q.Preload(relation)
	}

	if err := q.First(&twoFactorAuthSession, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &twoFactorAuthSession, nil
}

func (r *TwoFactorAuthSessionRepository) Save(twoFactorAuthSession *entities.TwoFactorAuthSession) error {
	return r.db.Save(twoFactorAuthSession).Error
}
