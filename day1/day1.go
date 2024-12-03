package day1

import (
	"advent_of_code_2024/helper"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1() {
	var array1 []int
	var array2 []int
	array1, array2 = createArrayFromFile()
	helper.HeapSort(array1)
	helper.HeapSort(array2)

	// Part 1
	var distance int = getDistance(array1, array2)
	fmt.Printf("Total distance is %d\n", distance)
	// Part 2
	var score int = getSimilarityScore(array1, array2)
	fmt.Printf("Total Score is %d\n", score)

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

	var total_distance int = 0
	for i := 0; i < len(array1); i++ {
		var distance = helper.Abs(array1[i] - array2[i])
		total_distance = total_distance + distance
		// fmt.Printf("i: %d, arr1: %d, arr2: %d, distance: %d\n", i, array1[i], array2[i], distance)
	}
	return total_distance
}

func getSimilarityScore(array1 []int, array2 []int) int {
	var array2copy = make([]int, len(array2))
	copy(array2copy, array2)

	var offset int = 0
	var totalScore int = 0
	for i := 0; i < len(array1); i++ {
		var count int = 0

		for j := 0; j < len(array2copy); j++ {
			if array1[i] == array2copy[j] {
				// fmt.Printf("%d %d\n", array1[i], array2[j])
				// Move the matching element to the offset position
				array2copy[offset], array2copy[j] = array2copy[j], array2copy[offset]
				offset++
				count++
			}
		}
		var score int = array1[i] * count
		totalScore += score
	}

	// for i := 0; i < len(array1); i++ {
	// 	fmt.Printf("%d\n", array2copy[i])
	// }
	return totalScore
}
