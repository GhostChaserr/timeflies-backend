package main

import (
	"fmt"
	"microservice_app/auth_service/src/handlers/healthhandler"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Registering routes
	router := mux.NewRouter()
	router.HandleFunc("/", healthhandler.HealthHandler)

	PORT := os.Getenv("PORT")

	fmt.Println("Running on -----------", PORT)
	http.ListenAndServe(":"+PORT, router)
}
