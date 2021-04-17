package middleware

import (
	"log"
	"net/http"

	"github.com/ebalkanski/middleware/internal/api"
)

type Cache interface {
	Message() string
}

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
			msg := api.NewErrorResponse("Content-Type header is missing")
			api.WriteResponse(w, msg, http.StatusBadRequest)
			return
		}
		if contentType != "application/json" {
			msg := api.NewErrorResponse("Content-Type header must be application/json")
			api.WriteResponse(w, msg, http.StatusBadRequest)
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

func Caching(logger *log.Logger, cache Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Executing middleware Caching")

			msg := cache.Message()
			if msg != "" {
				api.WriteResponse(w, api.NewSuccessResponse(msg), http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
