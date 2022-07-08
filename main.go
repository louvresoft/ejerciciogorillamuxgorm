package main

import (
	"github.com/gorilla/mux"
	"github.com/louvresoft/go-gorm-gorilla-mux/db"
	"github.com/louvresoft/go-gorm-gorilla-mux/models"
	"github.com/louvresoft/go-gorm-gorilla-mux/routes"
	"net/http"
)

func main() {
	db.DBConnection()
	err := db.DB.AutoMigrate(models.Task{})
	if err != nil {
		return
	}
	err = db.DB.AutoMigrate(models.User{})
	if err != nil {
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHAndler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
