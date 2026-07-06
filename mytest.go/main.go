package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("api/v1/",API)
	http.HandleFunc("/ping",ping)
	http.HandleFunc("/v2",v2Handler)
	http.HandleFunc("/dashboard",dashboard)
	http.HandleFunc("/userAgent",userAgent)
	http.HandleFunc("/cal",cal)
	http.HandleFunc("/handler", handler)
	fmt.Println("server is runing on 8080")
	http.ListenAndServe(":8080", nil)
}
