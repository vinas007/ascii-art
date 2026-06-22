package main

import "strings"

func render(input string, banner map[rune][]string) string {
	var out strings.Builder
	words := strings.Split(input, "\\n")
	for i, word := range words {
		if word == "" {
			if i < len(word) -1 {
				out.WriteByte('\n')
			}
		}
		for row := 0; row < 8; row++ {
			for _, char := range word {
				out.WriteString(banner [char][row])
			}
			out.WriteByte('\n')
		}
	}
	return out.String()
}