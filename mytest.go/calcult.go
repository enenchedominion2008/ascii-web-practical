package main

import (
	"net/http"
	"strconv"
	"fmt"
)

func cal(w http.ResponseWriter, r *http.Request) {
	add := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	s, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "not a valid number", http.StatusInternalServerError)
		return
	}
	m,err := strconv.Atoi(b)
	if err != nil {
		http.Error(w,"not a valid string",http.StatusInternalServerError)
		return
	}
	if add == "add" {
		result := s+m
		fmt.Fprintf(w,"result %d",result)
	} else if add == "sub"{
		result := s-m
		fmt.Fprintf(w,"result %d",result)
	} else if add == "mul"{
		result  := s*m
		fmt.Fprintf(w,"result %d",result)
		
	} else if add == "div" {
		if m == 0 {
			http.Error(w,"not divisable with zero",http.StatusBadRequest)
			return
		}
		result := s/m 
		fmt.Fprintf(w,"result %d",result)
		
	} else {
		http.Error(w,"not an conversion method",http.StatusBadRequest)
		return
	}

}
