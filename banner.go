package main 

import(
	"errors"
	"strings"
	"os"
)

func loadBanner(banner string) (map[rune][]string, error) {
	file, err := os.ReadFile(banner)
	if err != nil {
		return nil, errors.New("invalid banner file")
	}

	if len(file) == 0 {
		return nil, errors.New("empty banner file")
	}

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	if len(lines) != 856 {
		return nil, errors.New("incomplete banner file")
	}

	star := make(map[rune][]string)
	for i := rune(' '); i <= '~'; i++ {
		start := (i-' ') * 9 + 1
		star[i] = lines[start : start + 8]
	} 
	return star, nil
}