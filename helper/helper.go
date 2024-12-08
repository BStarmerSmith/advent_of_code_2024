package helper

import (
	"bufio"
	"log"
	"os"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MatrixFromFile(filePath string) [][]rune {
	var reports [][]rune

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		reports = append(reports, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return reports
}

// HeapSort sorts an array using the heap sort algorithm
func HeapSort(arr []int) {
	arrLen := len(arr)

	// Build the max heap
	for i := arrLen/2 - 1; i >= 0; i-- {
		heapify(arr, arrLen, i)
	}

	// Extract elements from the heap
	for i := arrLen - 1; i > 0; i-- {
		// Move the root (largest element) to the end
		arr[0], arr[i] = arr[i], arr[0]
		// Restore the heap property for the reduced heap
		heapify(arr, i, 0)
	}
}

// heapify ensures the heap property is maintained
func heapify(arr []int, size int, root int) {
	largest := root
	left := 2*root + 1  // Left child index
	right := 2*root + 2 // Right child index
	// Check if the left child is larger than the root
	if left < size && arr[left] > arr[largest] {
		largest = left
	}
	// Check if the right child is larger than the largest so far
	if right < size && arr[right] > arr[largest] {
		largest = right
	}
	// If the largest is not the root, swap and continue heapifying
	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, size, largest)
	}
}
