package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatobaez/gorm-resapi/application"
	"github.com/renatobaez/gorm-resapi/domain"
	"github.com/renatobaez/gorm-resapi/infrastructure/db"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(domain.User{}, domain.Task{})
	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler)

	r.HandleFunc("/users", application.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users", application.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", application.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", application.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/tasks", application.GetAllTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", application.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", application.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", application.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
