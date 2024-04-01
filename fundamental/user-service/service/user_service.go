package service

import (
	"github.com/ibrahimker/golang-praisindo/fundamental/user-service/entity"
)

type UserService struct {
	userDB  []entity.User
	userMap map[string]entity.User
}

func NewUserSvc(userDB []entity.User) *UserService {
	return &UserService{
		userDB: userDB,
	}
}

func (u *UserService) Register(user entity.User) (entity.User, error) {
	u.userDB = append(u.userDB, user)
	return user, nil
}

func (u *UserService) GetAll() ([]entity.User, error) {
	return u.userDB, nil
}

func (u *UserService) GetByEmail(email string) (entity.User, error) {
	return u.userMap[email], nil
}
