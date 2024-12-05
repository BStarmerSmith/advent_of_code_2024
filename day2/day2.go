package day2

import (
	"advent_of_code_2024/helper"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Day2 is a function that processes reports from a file and prints the number of good reports.
func Day2() {
	reports := readReportsFromFile("day2/input.txt")
	goodReports, badReports := processReports(reports)
	fmt.Printf("Good Reports: %d\nBad Reports: %d\n", goodReports, badReports)
}

// readReportsFromFile reads reports from a file and returns them as a 2D slice of integers.
func readReportsFromFile(filePath string) [][]int {
	var reports [][]int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels, err := parseLineToIntArray(line)
		if err != nil {
			log.Fatalf("Error parsing line: %v", err)
		}
		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return reports
}

// parseLineToIntArray parses a line of space-separated integers and returns them as a slice of integers.
func parseLineToIntArray(line string) ([]int, error) {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))

	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		levels[i] = value
	}
	return levels, nil
}

// processReports processes the reports and returns the number of good reports.
func processReports(reports [][]int) (int, int) {
	badReports := 0
	goodReports := 0
	for _, report := range reports {
		isIncrementing := report[0] <= report[1]

		if isValidReport(report, isIncrementing) {
			goodReports++
		} else {
			badReports++
		}
	}

	return goodReports, badReports
}

// isValidReport checks if a report is valid based on the given criteria.
func isValidReport(report []int, isIncrementing bool) bool {
	failover := false
	for i := 0; i < len(report)-1; i++ {
		if !isValidTransition(report[i], report[i+1], isIncrementing) {
			if !failover {
				failover = true
				isIncrementing = report[i] <= report[i+1]
				continue
			}
			return false
		}
	}
	return true
}

// isValidTransition checks if a transition between two levels is valid based on the given criteria.
func isValidTransition(current, next int, isIncrementing bool) bool {
	return checkIfIncrementing(current, next, isIncrementing) &&
		checkIfLevelsDiffer(current, next)
}

// checkIfLevelsDiffer checks if the difference between two levels is within the allowed range.
func checkIfLevelsDiffer(current, next int) bool {
	difference := helper.Abs(current - next)
	return difference > 0 && difference <= 3
}

// checkIfIncrementing checks if the levels are incrementing or decrementing based on the given criteria.
func checkIfIncrementing(current, next int, isIncrementing bool) bool {
	return (isIncrementing && current <= next) || (!isIncrementing && current >= next)
}
