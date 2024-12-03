package day1

import (
	"advent_of_code_2024/helper"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1Part1() {
	var array1 []int
	var array2 []int
	array1, array2 = createArrayFromFile()
	helper.HeapSort(array1)
	helper.HeapSort(array2)
	var distance int = getDistance(array1, array2)
	fmt.Printf("Total distance is %d", distance)
}

func createArrayFromFile() ([]int, []int) {
	// Use the relative path to input.txt
	filePath := "day1/input.txt"
	var array1 []int
	var array2 []int

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close() // Ensure the file is closed when the function exits

	scanner := bufio.NewScanner(file)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text() // Get the text of the current line
		inputs := strings.Split(line, "   ")
		array1value, err := strconv.Atoi(inputs[0])
		if err != nil {
			// ... handle error
			panic(err)
		}
		array2value, err := strconv.Atoi(inputs[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		array1 = append(array1, array1value)
		array2 = append(array2, array2value)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return array1, array2
}

func getDistance(array1 []int, array2 []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var total_distance int = 0
	for i := 0; i < len(array1); i++ {
		var distance = abs(array1[i] - array2[i])
		total_distance = total_distance + distance
		// fmt.Printf("i: %d, arr1: %d, arr2: %d, distance: %d\n", i, array1[i], array2[i], distance)
	}
	return total_distance
}
