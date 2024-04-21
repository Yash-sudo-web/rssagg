package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errJSON(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Error: %v", msg)
	}
	type errRes struct {
		Error string `json:"error"`
	}
	responseJSON(w, code, errRes{Error: msg})
}

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
