package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fdaygon/drift/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Welcome to Drift")
	router := chi.NewRouter()

	router.Get("/", handlers.GetDashboard)

	log.Fatal(http.ListenAndServe(":8080", router))
}
