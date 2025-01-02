package routes

import (
	"github.com/gorilla/mux"
	"github.com/rzashakh/todoro/handlers"
)

func SetupRoutes() *mux.Router { // Renamed to export the function
	router := mux.NewRouter()

	router.HandleFunc("/todos", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTask).Methods("DELETE")

	return router
}
