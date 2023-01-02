package service

import (
	"context"
	"e-learning/entity"
	"e-learning/repository"
	"e-learning/utils"
	"errors"
	"log"
	"time"
)

type AdminService interface {
	Login(ctx context.Context, admin *entity.Admin) (id int, err error)
	Register(ctx context.Context, admin *entity.Admin) (entity.Admin, error)
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) *adminService {
	return &adminService{adminRepo}
}

func (a *adminService) Login(ctx context.Context, admin *entity.Admin) (int, error) {
	dbAdmin, err := a.adminRepository.GetAdminByUsername(ctx, admin.Username)
	if err != nil {
		return 0, err
	}

	if dbAdmin.Username == "" || dbAdmin.ID == 0 {
		return 0, errors.New("admin not found")
	}

	password := utils.CheckPasswordHash(admin.Password, dbAdmin.Password)
	if !password {
		return 0, errors.New("wrong user or password")
	}

	return dbAdmin.ID, nil
}

func (a *adminService) Register(ctx context.Context, admin *entity.Admin) (entity.Admin, error) {
	dbAdmin, err := a.adminRepository.GetAdminByUsername(ctx, admin.Username)
	if err != nil {
		log.Println(err)
		return *admin, err
	}

	if dbAdmin.Username != "" || dbAdmin.ID != 0 {
		log.Println("Email Already")
		return *admin, errors.New("username already exits")
	}

	admin.CreatedAt = time.Now()
	newAdmin, err := a.adminRepository.CreateAdmin(ctx, *admin)
	if err != nil {
		log.Println(err)
		return *admin, err
	}

	return newAdmin, nil
}
