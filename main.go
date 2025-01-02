package main

import (
	"log"
	"net/http"

	"github.com/rzashakh/todoro/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
