package main


import (
	"net/http"
	"fmt"
)

func legacyHandler(w http.ResponseWriter, r *http.Request) {
http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}
func v2Handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Welcome to version 2")
}