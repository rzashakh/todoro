package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rzashakh/todoro/handlers"
)

func SetupRoutes() *mux.Router { // Renamed to export the function
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", handlers.MainPage).Methods("GET")
	router.HandleFunc("/todos", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTask).Methods("DELETE")

	return router
}
