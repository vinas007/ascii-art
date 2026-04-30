package main

import (
	"fmt"
)

func renderLine(line string, banner []string) {

	if line == "" {
		fmt.Println()
		return
	}
	for row := 1; row <= 8; row++ {
		for _, char := range line {
			if char < 32 || char > 126 {
				continue
			}
			index := (int(char) - 32) * 9
			fmt.Print(banner[index+row])
		}
		fmt.Println()
	}

}
