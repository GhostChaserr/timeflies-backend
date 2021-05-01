package healthhandler

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{"message": "user searvice"})
}
