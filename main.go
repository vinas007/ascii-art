package main

import(
	"strings"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . input [banner]")
		return
	}

	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		arg := os.Args[2]
		if !strings.HasSuffix(arg, ".txt" ) {
			arg = arg + ".txt"
		}
		bannerFile = arg
	}
	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(render(os.Args[1], banner))
}