package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var header = request.Header.Get("x-access-token")

		json.NewEncoder(response).Encode(request)
		header = strings.TrimSpace(header)

		if header == "" {
			response.WriteHeader(http.StatusForbidden)
			json.NewEncoder(response).Encode("Missing auth token")
			return
		} else {
			json.NewEncoder(response).Encode(fmt.Sprintf("Token found. Value %s", header))
		}

		ctx := context.WithValue(request.Context(), "user", "george")
		next.ServeHTTP(response, request.WithContext(ctx))
	})
}
