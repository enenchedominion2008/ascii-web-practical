package main


import (
	"net/http"
	"fmt"
)
func API(w http.ResponseWriter, r *http.Request) {
if r.URL.Path == "/api/v1/greet" {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "friend"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
    return
}
if r.URL.Path == "/api/v1/ping" {
    fmt.Fprintln(w, "pong")
    return
}

}