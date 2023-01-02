package service

import (
	"context"
	"e-learning/entity"
	"e-learning/repository"
	"e-learning/utils"
	"errors"
	"time"
)

type UserService interface {
	Login(ctx context.Context, user *entity.User) (id int, err error)
	Register(ctx context.Context, user *entity.User) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	taskRepository repository.TaskRepository
}

func NewUserService(userRepository repository.UserRepository, taskRepo repository.TaskRepository) UserService {
	return &userService{userRepository, taskRepo}
}

func (u *userService) Login(ctx context.Context, user *entity.User) (id int, err error) {
	dbUser, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return 0, err
	}

	if dbUser.Email == "" || dbUser.ID == 0 {
		return 0, errors.New("user not found")
	}

	// cek Password
	password := utils.CheckPasswordHash(user.Password, dbUser.Password)

	if !password {
		return 0, errors.New("wrong User or Password")
	}

	return dbUser.ID, nil
}

func (u *userService) Register(ctx context.Context, user *entity.User) (entity.User, error) {
	dbUser, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exits")
	}

	user.CreatedAt = time.Now()

	newUser, err := u.userRepository.CreateUser(ctx, *user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}
