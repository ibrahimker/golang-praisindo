package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/golang-praisindo/user-service-http/entity"
	"github.com/ibrahimker/golang-praisindo/user-service-http/service"
)

type UserHandler struct {
	userServices service.IUserService
}

type IUserHandler interface {
	GetAllUsersHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)

	GetUserByEmailHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserByEmailHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserByEmailHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userServices service.IUserService) IUserHandler {
	return &UserHandler{
		userServices: userServices,
	}
}

func (u *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := u.userServices.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// print user if success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// read json data
	userRequest := entity.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create data
	user, err := u.userServices.Create(userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// print user if success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	user, err := u.userServices.GetByEmail(mux.Vars(r)["email"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// print user if success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) UpdateUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	// read json data
	userRequest := entity.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRequest.Email = mux.Vars(r)["email"]

	// update data in service
	user, err := u.userServices.Update(userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// print user if success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) DeleteUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	err := u.userServices.Delete(mux.Vars(r)["email"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successfully delete user")
}
