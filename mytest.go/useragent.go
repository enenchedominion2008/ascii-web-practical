package main

import (
	"fmt"
	"net/http"
)

func userAgent(w http.ResponseWriter, r *http.Request) {
	UserAgent := r.Header.Get("UserAgent")
	if UserAgent == ""{
		UserAgent =  "unknown device"
	}

	fmt.Fprintf(w,"you are visiting us using %s",UserAgent)
}