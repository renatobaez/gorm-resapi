package application

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatobaez/gorm-resapi/domain"
	"github.com/renatobaez/gorm-resapi/infrastructure/db"
)

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []domain.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdTask.Error.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	db.DB.First(&task, mux.Vars(r)["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	db.DB.First(&task, mux.Vars(r)["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}
