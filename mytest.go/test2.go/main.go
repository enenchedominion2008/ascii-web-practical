package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/method-inspector", Method)
	http.HandleFunc("/echo",Echo)
	fmt.Println("server is runing on 8080")
	http.ListenAndServe(":8080", nil)
}
