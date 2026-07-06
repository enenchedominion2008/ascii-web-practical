package main

import (
	"io"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}
	data,err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"could nit read body",http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if len(data) == 0 {
		http.Error(w,"body can not be empty",http.StatusBadRequest)
		return
	}
	w.Header().Set("content-Type","text/Plain")
	w.Write(data)
	
}
