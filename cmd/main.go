package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ebalkanski/go-middleware/internal/service"
)

func main() {
	svc := service.New()

	r := mux.NewRouter()
	r.HandleFunc("/", svc.Hello)

	// configure http server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("server error")
	}
}
