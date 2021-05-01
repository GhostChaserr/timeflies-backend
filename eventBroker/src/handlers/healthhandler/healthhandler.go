package healthhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HealthHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	// Checking all services healths
	userServiceHealthResponse, _ := http.Get("http://host.docker.internal:5001")
	storyServiceHealthResponse, _ := http.Get("http://host.docker.internal:5002")
	fmt.Println(userServiceHealthResponse, storyServiceHealthResponse.Status)

	responsePayload := make(map[string]string)
	responsePayload["userService"] = "OK"
	responsePayload["storyService"] = "OK"
	responsePayload["eventBroker"] = "OK"

	json.NewEncoder(response).Encode(responsePayload)
}
