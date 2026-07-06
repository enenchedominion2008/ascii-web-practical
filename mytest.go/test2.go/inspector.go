package main

import (
	"fmt"
	"net/http"
)

func Method(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("you made a %s request",r.Method)
	fmt.Fprintf(w,msg)
	
	

}
