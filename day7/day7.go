package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var symbolMap []string

func init() {
	symbolMap = []string{"+", "*"}
}

func Day7() {
	inputFileString := "day7/input.txt"
	equationMap := generateEquationMap(inputFileString)
	log.Printf("Equation map: %v", equationMap)
	totalCount := 0
	for key, value := range equationMap {
		count := calculateEquation(key, value)
		totalCount += count
		log.Printf("Key: %d, Value: %d, Count: %d", key, value, count)
	}

	log.Printf("Total count: %d", totalCount)
}

func calculateEquation(total int, numbers []int) int {
	count := 0
	n := len(numbers)
	operationCount := len(symbolMap)

	// Generate all possible combinations of operations
	// There are (operationCount^(n-1)) combinations
	combinationCount := 1
	for i := 0; i < n-1; i++ {
		combinationCount *= operationCount
	}

	for i := 0; i < combinationCount; i++ {
		ops := generateOperations(i, n-1)
		if evaluateEquation(numbers, ops) == total {
			count++
		}
	}
	return count
}

// Helper to generate a sequence of operations based on an index
func generateOperations(index, length int) []string {
	ops := make([]string, length)
	for i := 0; i < length; i++ {
		ops[i] = symbolMap[index%len(symbolMap)]
		index /= len(symbolMap)
	}
	return ops
}

// Helper to evaluate a sequence of numbers and operations
func evaluateEquation(numbers []int, operations []string) int {
	result := numbers[0]
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		}
	}
	return result
}

// Exaple input:
// 479027832: 8 9 69 659 96 634
// 539373: 6 19 6 863
// 50830: 27 9 91 25 2
func generateEquationMap(inputFileString string) map[int][]int {
	equationMap := make(map[int][]int)

	file, err := os.Open(inputFileString)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				log.Fatalf("Error converting key to int: %v", err)
			}
			values := strings.Fields(strings.TrimSpace(parts[1]))
			intValues := make([]int, len(values))
			for i, v := range values {
				intValue, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalf("Error converting value to int: %v", err)
				}
				intValues[i] = intValue
			}
			equationMap[key] = intValues
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return equationMap
}
