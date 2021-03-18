package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/ebalkanski/go-middleware/internal/service"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		next.ServeHTTP(w, r)
	})
}

func main() {
	svc := service.New()
	chain := alice.New(middlewareOne, middlewareTwo).ThenFunc(svc.Hello)
	r := mux.NewRouter()
	r.
		Methods("GET").
		Path("/").
		Handler(chain)

	// configure http server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("server error")
	}
}
