package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3() {
	programRaw := readProgramFromFile("day3/input.txt")
	cleanData := parseProgramData(programRaw)
	total := processCalculation(cleanData)
	fmt.Printf("total: %d\n", total)
}

// readReportsFromFile reads reports from a file and returns them as a 2D slice of integers.
func readProgramFromFile(filePath string) []string {
	var reports []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, line)
	}
	return reports
}

func parseProgramData(rawDataArr []string) []string {
	var cleanData []string
	for _, rawData := range rawDataArr {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllString(rawData, -1)
		cleanData = append(cleanData, matches...) // ... is used to expand the values
	}
	return cleanData
}

func processCalculation(cleanData []string) int {
	total := 0
	for _, data := range cleanData {
		re := regexp.MustCompile(`\b\d{1,3}\b`)
		matches := re.FindAllString(data, -1)
		value1 := parseStringToInt(matches[0])
		value2 := parseStringToInt(matches[1])
		total += (value1 * value2)
	}
	return total
}

func parseStringToInt(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	return value
}
