package main

import (
	"fmt"

	"github.com/ibrahimker/golang-praisindo/fundamental/user-service/entity"
	"github.com/ibrahimker/golang-praisindo/fundamental/user-service/service"
)

func main() {
	userDB := make([]entity.User, 0)
	userSvc := service.NewUserSvc(userDB) // contoh dependency injection

	if user, err := userSvc.Register(entity.User{
		Username: "budi123",
		Email:    "budi123@gmail.com",
		Password: "password123",
		Age:      9,
	}); err != nil {
		fmt.Printf("Error when register user: %+v\n", err)
		return
	} else {
		fmt.Printf("Success register user: %+v\n", user)
	}

	res, _ := userSvc.GetAll()
	fmt.Printf("GetAll results: %+v\n", res)

}
