package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatobaez/gorm-resapi/internal/application"
	"github.com/renatobaez/gorm-resapi/internal/domain"
)

type UserController struct {
	userService application.UserService
}

func NewUserController(userService application.UserService) *UserController {
	return &UserController{userService}
}

func (uc *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := uc.userService.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&res)
}
func (uc *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	res, err := uc.userService.GetUser(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&res)
}
func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	if err := uc.userService.CreateUser(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (uc *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := uc.userService.DeleteUser(mux.Vars(r)["id"]); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
