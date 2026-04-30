package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		return
	}

	input := os.Args[1]

	input = strings.ReplaceAll(input, "\\n", "\n")

	bannerFile := getBannerFile(os.Args)
	if bannerFile == "" {
		return
	}

	banner := loadBanner(bannerFile)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		renderLine(line, banner)
	}
}
