package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewSuccessResponse(msg string) *Response {
	return &Response{
		Message: msg,
	}
}

func NewErrorResponse(err string) *Response {
	return &Response{
		Error: err,
	}
}

func WriteResponse(w http.ResponseWriter, msg *Response, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Println("Cannot write response")
	}
}
