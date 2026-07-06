package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring> \"something\"")
		return
	}

	if !strings.HasPrefix(os.Args[1],"--color=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring> \"something\"")
		return
	}
	colorName := strings.TrimPrefix(os.Args[1],"--color=")
	substring := os.Args[2] 
	text := os.Args[3]
	fmt.Println("color:",colorName)
	fmt.Println("substring:",substring)
	fmt.Println("text:",text)
}