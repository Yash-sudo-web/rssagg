package main

import "net/http"

func handle_err(w http.ResponseWriter, r *http.Request) {
	errJSON(w, 500, "This is an error")
}
