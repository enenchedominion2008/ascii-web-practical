package main

import (
	"fmt"
	"net/http"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	dashboard := r.Header.Get("X-API-Key")

	if dashboard != "secrect123" {
		http.Error(w, "not acepted you are blocked", http.StatusUnauthorized)
	} else {
		fmt.Fprintf(w, "welcome to the dashboard")
	}
}
