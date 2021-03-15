package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"middleware/internal/service"
	"net/http"
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
