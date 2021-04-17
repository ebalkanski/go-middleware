package service

import (
	"log"
	"net/http"

	"github.com/ebalkanski/middleware/internal/api"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Service.Hello() is called")

	msg := api.NewSuccessResponse("Hello")
	api.WriteResponse(w, msg, http.StatusOK)
}
