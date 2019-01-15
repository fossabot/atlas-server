package main

import (
	"github.com/atlas-io/atlas-server/controllers/projects"
	"github.com/atlas-io/atlas-server/controllers/tasks"
	"github.com/gorilla/mux"
)

// SetupRoutes sets up all API endpoints
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/projects", projects.PutProjectRoute).Methods("PUT")
	r.HandleFunc("/projects/{key}", projects.GetProjectRoute).Methods("GET")

	r.HandleFunc("/tasks", tasks.PutTaskRoute).Methods("PUT")
	r.HandleFunc("/tasks/{id}", tasks.GetTaskRoute).Methods("GET")
}
