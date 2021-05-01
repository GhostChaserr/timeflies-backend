package main

import (
	"fmt"
	"microservice_app/story_service/src/controllers/storycontroller"
	"microservice_app/story_service/src/handlers/healthhandler"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Registering routes
	router := mux.NewRouter()
	router.HandleFunc("/", healthhandler.HealthHandler)
	router.HandleFunc("/api/stories", storycontroller.GetStories()).Methods("GET")
	router.HandleFunc("/api/stories/{storyId}", storycontroller.GetStory()).Methods("GET")
	router.HandleFunc("/api/stories", storycontroller.CreateStory()).Methods("POST")

	PORT := os.Getenv("PORT")

	fmt.Println("Running on -----------", PORT)
	http.ListenAndServe(":"+PORT, router)
}
