package healthhandler

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	responsePayload := make(map[string]string)
	responsePayload["authService"] = "OK"

	json.NewEncoder(response).Encode(responsePayload)
}
