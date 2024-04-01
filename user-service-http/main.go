package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/golang-praisindo/user-service-http/entity"
	"github.com/ibrahimker/golang-praisindo/user-service-http/handler"
	"github.com/ibrahimker/golang-praisindo/user-service-http/repository"
	"github.com/ibrahimker/golang-praisindo/user-service-http/service"
)

func main() {
	r := mux.NewRouter()
	usersDB := []entity.User{}

	userRepository := repository.NewUserRepository(usersDB)
	userServices := service.NewUserSvc(userRepository)
	userHandler := handler.NewUserHandler(userServices)

	r.HandleFunc("/users", userHandler.UsersHandler)
	r.HandleFunc("/users/{email}", userHandler.UserHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Serving http at port 8080")
	log.Fatal(srv.ListenAndServe())
}
