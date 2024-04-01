package repository

import (
	"errors"

	"github.com/ibrahimker/golang-praisindo/user-service-http/entity"
)

type UserRepository struct {
	userDB []entity.User
}

type IUserRepository interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	GetAll() ([]entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Delete(email string) error
}

func NewUserRepository(userDB []entity.User) IUserRepository {
	return &UserRepository{
		userDB: userDB,
	}
}

func (u *UserRepository) GetAll() ([]entity.User, error) {
	return u.userDB, nil
}

func (u *UserRepository) GetByEmail(email string) (entity.User, error) {
	for _, user := range u.userDB {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

// Create implements IUserRepository.
func (u *UserRepository) Create(user entity.User) (entity.User, error) {
	u.userDB = append(u.userDB, user)
	return user, nil
}

// Delete implements IUserRepository.
func (u *UserRepository) Delete(email string) error {
	// copy all element to new slice except element that want to be deleted
	var userCopy []entity.User
	for _, user := range u.userDB {
		if user.Email != email {
			userCopy = append(userCopy, user)
		}
	}
	u.userDB = userCopy

	return nil
}

// Update implements IUserRepository.
func (u *UserRepository) Update(user entity.User) (entity.User, error) {
	// naive approach to update is by delete the data first then create the data again
	if err := u.Delete(user.Email); err != nil {
		return entity.User{}, err
	}

	newUser, err := u.Create(user)
	if err != nil {
		return entity.User{}, err
	}

	return newUser, nil
}
