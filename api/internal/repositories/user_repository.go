package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{
		db:  tx,
		ctx: r.ctx,
	}
}

func (r *UserRepository) WithContext(ctx context.Context) *UserRepository {
	return &UserRepository{
		db:  r.db,
		ctx: ctx,
	}
}

func (r *UserRepository) Save(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) FindOneById(id uuid.UUID) (*entities.User, error) {
	var user entities.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindOneByUsernameUnscoped(username string, excludedIds ...uuid.UUIDs) (*entities.User, error) {
	var user entities.User

	q := r.db.Unscoped().Where("username = ?", username)

	if len(excludedIds) > 0 {
		q = q.Where("id NOT IN ?", excludedIds)
	}

	if err := q.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindOneByEmailUnscoped(email string, excludedIds ...uuid.UUIDs) (*entities.User, error) {
	var user entities.User

	q := r.db.Unscoped().Where("email = ?", email)

	if len(excludedIds) > 0 {
		q = q.Where("id NOT IN ?", excludedIds)
	}

	if err := q.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindOneByUsernameOrEmail(usernameOrEmail string) (*entities.User, error) {
	var user entities.User

	if err := r.db.Preload("UserDevices").First(&user, "username = ? OR email = ?", usernameOrEmail, usernameOrEmail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
