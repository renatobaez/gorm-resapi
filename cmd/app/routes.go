package app

import "github.com/gorilla/mux"

func LoadRoutes(r *mux.Router, app *Handlers) {

	r.HandleFunc("/users", app.userController.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users", app.userController.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", app.userController.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", app.userController.DeleteUserHandler).Methods("DELETE")

	// r.HandleFunc("/tasks", application.GetAllTasksHandler).Methods("GET")
	// r.HandleFunc("/tasks", application.CreateTaskHandler).Methods("POST")
	// r.HandleFunc("/tasks/{id}", application.GetTaskHandler).Methods("GET")
	// r.HandleFunc("/tasks/{id}", application.DeleteTaskHandler).Methods("DELETE")
}
