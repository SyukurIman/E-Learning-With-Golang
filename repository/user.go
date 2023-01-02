package repository

import (
	"context"
	"e-learning/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserById(ctx context.Context, id int) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	data := entity.User{}
	err := r.db.Model(&data).Where("email = ?", email).Scan(&data).Error
	return data, err
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (entity.User, error) {
	data := entity.User{}
	err := r.db.Model(&data).Where("id = ?", id).Scan(&data).Error
	return data, err
}
