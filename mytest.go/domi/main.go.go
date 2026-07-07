package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring> \"something\"")
		return
	}

	if !strings.HasPrefix(os.Args[1], "--color=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring> \"something\"")
		return
	}

	colorName := strings.TrimPrefix(os.Args[1], "--color=")

	var substring, text, banner string
	banner = "standard.txt"

	if len(os.Args) == 3 {
		// go run . --color=red "hello"
		text = os.Args[2]
		substring = os.Args[2]

	} else if len(os.Args) == 4 {
		// go run . --color=red "hello" shadow
		// OR: go run . --color=red kit "hello"
		if os.Args[3] == "shadow" || os.Args[3] == "thinkertoy" {
			text = os.Args[2]
			substring = os.Args[2]
			banner = os.Args[3] + ".txt"
		} else {
			substring = os.Args[2]
			text = os.Args[3]
		}

	} else if len(os.Args) == 5 {
		// go run . --color=red kit "hello" shadow
		substring = os.Args[2]
		text = os.Args[3]
		banner = os.Args[4] + ".txt"
	}

	color := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"purple":  "\033[35m",
		"skyblue": "\033[36m",
		"white":   "\033[37m",
	}

	ansiCode, exist := color[colorName]
	if !exist {
		fmt.Println("error: color \"" + colorName + "\" is not supported")
		return
	}

	result := ascii(text, banner, substring, ansiCode)
	fmt.Print(result)
}