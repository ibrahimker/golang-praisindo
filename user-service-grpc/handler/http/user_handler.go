package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/entity"
	"github.com/ibrahimker/golang-praisindo/user-service-grpc/service"
)

type UserHandler struct {
	userServices service.IUserService
}

type IUserHandler interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
	UserHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userServices service.IUserService) IUserHandler {
	return &UserHandler{
		userServices: userServices,
	}
}

// UsersHandler implements IUserHandler.
func (u *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		u.createUserHandler(w, r)
	case http.MethodGet:
		u.getAllUsersHandler(w, r)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}

func (u *UserHandler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := u.userServices.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// print user if success
	json.NewEncoder(w).Encode(users)
}

func (u *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(user)
}

// UserHandler implements IUserHandler.
func (u *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		u.getUserByEmailHandler(w, r)
	case http.MethodPut:
		u.updateUserByEmailHandler(w, r)
	case http.MethodDelete:
		u.deleteUserByEmailHandler(w, r)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}

func (u *UserHandler) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	user, err := u.userServices.GetByEmail(mux.Vars(r)["email"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// print user if success
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) updateUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) deleteUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	err := u.userServices.Delete(mux.Vars(r)["email"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successfully delete user")
}
