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
	color := map[string]string{
		"red" : "\033[31m",
		"green" : "\033[32m",
		"yellow" : "\033[33m",
		"blue" : "\033[34m",
		"purple" : "\033[35m",
		"skyblue" : "\033[36m",
		"white" : "\033[37m",

	}
	ansiCode ,exist := color[colorName] 
	if !exist {
		fmt.Println("error: color \""+colorName+"\"is not surported")
		return
	}
	fmt.Println("ANSI code:",ansiCode)
	result := ascii(text, "standard.txt")
fmt.Print(result)
}


