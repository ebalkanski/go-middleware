package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ebalkanski/go-middleware/internal/api"
)

func ResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware ResponseHeaders ")

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func RequestHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware RequestHeaders")

		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			w.WriteHeader(http.StatusBadRequest)
			msg := &api.Response{
				Error: "Content-Type header is missing",
			}
			if err := json.NewEncoder(w).Encode(msg); err != nil {
				log.Println("Cannot write response")
			}
			return
		}
		if contentType != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			msg := &api.Response{
				Error: "Content-Type header must be application/json",
			}
			if err := json.NewEncoder(w).Encode(msg); err != nil {
				log.Println("Cannot write response")
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Executing middleware Logging")

			next.ServeHTTP(w, r)
		})
	}
}
