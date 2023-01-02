package repository

import (
	"context"
	"e-learning/entity"

	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(ctx context.Context, admin entity.Admin) (entity.Admin, error)
	GetAdminByUsername(ctx context.Context, username string) (entity.Admin, error)
}
type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) CreateAdmin(ctx context.Context, admin entity.Admin) (entity.Admin, error) {
	return admin, r.db.Create(&admin).Error
}

func (r *adminRepository) GetAdminByUsername(ctx context.Context, username string) (entity.Admin, error) {
	data := entity.Admin{}
	return data, r.db.Model(&data).Where("username = ?", username).Scan(&data).Error
}
