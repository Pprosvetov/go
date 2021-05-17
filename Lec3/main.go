package main

import "fmt"

func main() {

	var (
		height int
		width  int
	)
	fmt.Scan(&height, &width)

	var per, sq = (height * 2) + (width * 2), height * width

	fmt.Printf("Периметр: %d\nПлощадь: %d\n", per, sq)
}
