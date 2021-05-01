package main

import (
	"fmt"
	"microservice_app/event_broker/src/handlers/healthhandler"
	"microservice_app/event_broker/src/handlers/storyhandler"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Registering routes
	router := mux.NewRouter()
	router.HandleFunc("/", healthhandler.HealthHandler)
	router.HandleFunc("/api/stories", storyhandler.CreateStory()).Methods("POST")

	PORT := os.Getenv("PORT")

	fmt.Println("Running on -----------", PORT)
	http.ListenAndServe(":"+PORT, router)
}
