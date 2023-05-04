package utils

import (
	"encoding/json"
	"net/http"
)

func BasicAuth(expectedUser, expectedPassword string) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, password, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(Map{"message": "forbidden"})
				return
			}
			if user != expectedUser || password != expectedPassword {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(Map{"message": "forbidden"})
				return
			}
			handler.ServeHTTP(w, r)
		})
	}
}
