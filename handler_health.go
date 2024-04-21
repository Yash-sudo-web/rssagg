package main

import "net/http"

func health(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, struct{}{})
}
