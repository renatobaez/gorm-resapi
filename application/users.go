package application

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatobaez/gorm-resapi/domain"
	"github.com/renatobaez/gorm-resapi/infrastructure/db"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []domain.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	db.DB.First(&user, mux.Vars(r)["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdUser.Error.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	db.DB.First(&user, mux.Vars(r)["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
