package storyhander

import (
	"encoding/json"
	"net/http"
)

func GetStories() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{"message": "fetching stories"})
	}
}

func CreateStory() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{"message": "creating stories"})
	}
}
