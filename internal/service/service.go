package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ebalkanski/go-middleware/internal/api"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Service.Hello() is called")

	var err error

	msg := &api.Response{
		Message: "Hello",
	}
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		fmt.Println("Cannot write response")
	}
}
