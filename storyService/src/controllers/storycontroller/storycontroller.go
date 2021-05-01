package storycontroller

import (
	"encoding/json"
	"fmt"
	"microservice_app/story_service/src/models/story"
	"microservice_app/story_service/src/payloads/createstory"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var stories = make([]story.Story, 0)
var id int64 = 0

func GetStories() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-type", "application/json")
		json.NewEncoder(response).Encode(stories)
	}
}

func GetStory() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		params := mux.Vars(request)
		storyId := params["storyId"]

		convertedStoryId, _ := strconv.ParseInt(storyId, 10, 64)
		for _, story := range stories {
			if story.Id == convertedStoryId {
				response.Header().Set("Content-type", "application/json")
				json.NewEncoder(response).Encode(story)
				return
			}
		}

		fmt.Println(storyId)
		response.Header().Set("Content-type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{"message": "getting story!"})
	}
}

func CreateStory() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var storyPayload createstory.CreateStory

		err := json.NewDecoder(request.Body).Decode(&storyPayload)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO. validations
		if len(storyPayload.Title) < 10 {
			http.Error(response, "Invalid title", 400)
			return
		}

		var story story.Story
		story.Id = id
		story.Media = storyPayload.Media
		story.Mood = storyPayload.Mood
		story.Summary = storyPayload.Summary
		story.Title = storyPayload.Title

		stories = append(stories, story)

		id += 1

		response.Header().Set("Content-type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{"message": "creating stories"})
	}
}
