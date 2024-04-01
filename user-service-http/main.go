package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/golang-praisindo/user-service-http/handler"
	"github.com/ibrahimker/golang-praisindo/user-service-http/repository"
	"github.com/ibrahimker/golang-praisindo/user-service-http/service"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	r := mux.NewRouter()

	// init database connection
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// setup repo-service-handler
	userRepository := repository.NewUserRepository(db)
	userServices := service.NewUserSvc(userRepository)
	userHandler := handler.NewUserHandler(userServices)

	// setup router
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
