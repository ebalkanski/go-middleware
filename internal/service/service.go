package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {}

func New() *Service {
	return &Service{}
}

type Response struct {
	Message string `json:"message"`
}

func (s *Service) Hello(w http.ResponseWriter, r *http.Request) {
	var err error

	msg := &Response{
		Message: "Hello",
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		fmt.Println("Cannot write response")
	}
}