package main

import (
	github.com/gorilla/mux
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}