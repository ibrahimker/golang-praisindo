package service

import (
	"errors"

	"github.com/ibrahimker/golang-praisindo/user-service-http/entity"
	"github.com/ibrahimker/golang-praisindo/user-service-http/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

type IUserService interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	GetAll() ([]entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Delete(email string) error
}

func NewUserSvc(repo repository.IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetAll() ([]entity.User, error) {
	return u.repo.GetAll()
}

func (u *UserService) GetByEmail(email string) (entity.User, error) {
	// input validation
	if email == "" {
		return entity.User{}, errors.New("email field is mandatory")
	}

	return u.repo.GetByEmail(email)
}

// Create implements IUserService.
func (u *UserService) Create(user entity.User) (entity.User, error) {
	// input validation
	if user.Email == "" {
		return entity.User{}, errors.New("email field is mandatory")
	}

	return u.repo.Create(user)
}

// Delete implements IUserService.
func (u *UserService) Delete(email string) error {
	// input validation
	if email == "" {
		return errors.New("email field is mandatory")
	}

	return u.repo.Delete(email)
}

// Update implements IUserService.
func (u *UserService) Update(user entity.User) (entity.User, error) {
	// input validation
	if user.Email == "" {
		return entity.User{}, errors.New("email field is mandatory")
	}

	return u.repo.Update(user)
}
