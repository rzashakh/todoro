package main

import (
	"github.com/rzashakh/todoro/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes() // Updated function call
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router)) // Start the server
}
