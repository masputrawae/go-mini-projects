package main

import (
	"fmt"
	"strings"
)

func main() {
	fruits := []string{
		"mangga", "mangga", "jeruk", "apel", "jeruk", "durian",
		"jeruk", "jeruk", "jeruk", "apel", "apel",
		"alpukat", "alpukat", "mangga", "jeruk", "durian",
	}

	total := len(fruits)
	countMap := make(map[string]int)

	for _, f := range fruits {
		countMap[f]++
	}

	for k, v := range countMap {
		p := Calculate(float64(v), float64(total))
		fmt.Printf("%s:\n", k)
		fmt.Printf("%s (%.2f%%)\n", strings.Repeat("#", int(p)), p)
	}
}

func Calculate(a, b float64) float64 {
	return (a / b) * 100
}
