package main

import (
	"fmt"
	"microservice_app/user_service/src/handlers/healthhandler"
	"microservice_app/user_service/src/handlers/storyhandler"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// Registering routes
	router := mux.NewRouter()
	router.HandleFunc("/", healthhandler.HealthHandler)
	router.HandleFunc("/api/events/stories", storyhandler.CreateStory()).Methods("POST")

	// db := connection.InitializeConnection()
	// err := connection.InsertUser(db, 2)

	// fmt.Println(db)
	PORT := os.Getenv("PORT")

	// defer db.Close()

	fmt.Println("Running on -----------", PORT)
	http.ListenAndServe(":"+PORT, router)
}
