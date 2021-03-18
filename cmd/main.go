package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/ebalkanski/go-middleware/internal/middleware"
	"github.com/ebalkanski/go-middleware/internal/service"
)

func main() {
	logger := log.Default()
	svc := service.New()
	chain := alice.New(middleware.ResponseHeaders, middleware.RequestHeaders, middleware.Logging(logger)).ThenFunc(svc.Hello)
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
