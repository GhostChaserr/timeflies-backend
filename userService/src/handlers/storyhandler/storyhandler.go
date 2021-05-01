package storyhandler

import (
	"encoding/json"
	"fmt"
	"microservice_app/user_service/src/payloads/story/publishstorypayload"
	"net/http"
)

func CreateStory() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		var createdStory publishstorypayload.CreateStory
		err := json.NewDecoder(request.Body).Decode(&createdStory)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Incoming event from event broker....")
		json.NewEncoder(response).Encode(map[string]string{"message": "Incoming storie"})

	}
}
