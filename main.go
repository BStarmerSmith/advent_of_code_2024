package main

import (
	"advent_of_code_2024/day1"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]

	switch day {
	case "1":
		day1.Day1Part1()
	// Add more cases as you implement more days
	default:
		fmt.Printf("Day %s is not implemented yet.\n", day)
	}
}