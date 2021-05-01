package storyhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"microservice_app/event_broker/src/constants"
	"microservice_app/event_broker/src/payloads/story/publishstorypayload"
	"net/http"
)

func CreateStory() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var storyPayload publishstorypayload.CreateStory

		err := json.NewDecoder(request.Body).Decode(&storyPayload)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		storyJson, err := json.Marshal(storyPayload)
		resp, err := http.Post(constants.UserServiceURI+"/api/events/stories", "application/json", bytes.NewBuffer(storyJson))

		if err != nil {
			http.Error(response, "Failed to send story payload", 500)
		}

		fmt.Println(resp.Status)

		response.Header().Set("Content-type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{"message": "new story from story service"})

	}
}
