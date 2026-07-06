package main

import (
	"os"
	"strings"
)

func ascii(text string, bannar string) string {
	data, err := os.ReadFile(bannar)
	if err != nil {

		return "could not readfile"
	}
	if len(data) == 0 {
		return "error banner is empty"
	}

	line := strings.Split(string(data), "\n")
	if len(line) != 855 {
		return "error banner is incomplect"
	}
	font := make(map[rune][]string)
	for char := rune(32); char <= rune(126); char++ {
		begin := (int(char) - 32) * 9
		font[char] = line[begin+1 : begin+9]
	}
	var final strings.Builder
	input := strings.ReplaceAll(text, "\\n", "\n")
	word := strings.Split(input, "\n")
	for i, j := range word {
		if j == "" {
			if i != len(word)-1 {
				final.WriteString("\n")
			}
			continue
			
		}
		for row := 0; row < 8; row++ {
			for _, char := range j {
			
				final.WriteString(font[char][row])
				
			}
			
			final.WriteString("\n")
		}

	}
	return final.String()

}
