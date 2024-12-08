package main

import (
	"advent_of_code_2024/day1"
	"advent_of_code_2024/day2"
	"advent_of_code_2024/day3"
	"advent_of_code_2024/day4"
	"advent_of_code_2024/day6"
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
		day1.Day1()
	case "2":
		day2.Day2()
	case "3":
		day3.Day3()
	case "4":
		day4.Day4()
	case "6":
		day6.Day6()
	// Add more cases as you implement more days
	default:
		fmt.Printf("Day %s is not implemented yet.\n", day)
	}
}
