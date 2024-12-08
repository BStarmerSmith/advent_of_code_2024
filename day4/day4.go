package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Diagonal down-right
	{-1, -1}, // Diagonal up-left
	{1, -1},  // Diagonal down-left
	{-1, 1},  // Diagonal up-right
}

func Day4() {
	words := []string{"XMAS"}
	programRaw := matrixFromFile("day4/input.txt")
	wordMap := findWordsInWordSearch(programRaw, words)
	for word := range wordMap {
		fmt.Printf("Word: %s Count: %d\n", word, len(wordMap[word]))
	}
	xMasPattern := findXMasPattern(programRaw)
	fmt.Printf("X-MAS Pattern Count: %d\n", len(xMasPattern))
}

func matrixFromFile(filePath string) [][]rune {
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

func findWordsInWordSearch(grid [][]rune, words []string) map[string][][3]int {
	results := make(map[string][][3]int)

	// Helper function to check if a word exists starting at a given position
	checkWord := func(grid [][]rune, word string, startRow, startCol int, direction Direction) bool {
		rows := len(grid)
		cols := len(grid[0])
		for i, char := range word {
			newRow := startRow + direction.dx*i
			newCol := startCol + direction.dy*i

			// Out of bounds or mismatch
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				return false
			}
			if grid[newRow][newCol] != char {
				return false
			}
		}
		return true
	}

	// Search through every cell in the grid for all words
	for _, word := range words {
		results[word] = [][3]int{} // Initialize with an empty list of positions
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[0]); col++ {
				if grid[row][col] == rune(word[0]) { // Potential start of the word
					for _, direction := range directions {
						if checkWord(grid, word, row, col, direction) {
							// Append start position and direction as a tuple
							results[word] = append(results[word], [3]int{row, col, directionIndex(direction)})
						}
					}
				}
			}
		}
	}

	return results
}

// Helper to get a unique identifier for each direction
func directionIndex(dir Direction) int {
	for i, d := range directions {
		if d == dir {
			return i
		}
	}
	return -1
}

func findXMasPattern(grid [][]rune) [][2]int {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])
	if cols == 0 {
		return nil
	}

	var results [][2]int

	// Helper to check if a diagonal matches any permutation of "MAS"
	isDiagonalXMas := func(diagonal []rune) bool {
		if len(diagonal) != 3 {
			return false
		}
		// Create a map to count occurrences of 'M', 'A', 'S'
		counts := make(map[rune]int)
		for _, char := range diagonal {
			counts[char]++
		}
		// "MAS" must have exactly 1 'M', 1 'A', 1 'S'
		return counts['M'] == 1 && counts['A'] == 1 && counts['S'] == 1
	}

	// Check for "X-MAS" pattern centered at (r, c)
	checkXMas := func(r, c int) bool {
		if r-1 >= 0 && r+1 < rows && c-1 >= 0 && c+1 < cols {
			// Extract the diagonals
			topLeftToBottomRight := []rune{grid[r-1][c-1], grid[r][c], grid[r+1][c+1]}
			topRightToBottomLeft := []rune{grid[r-1][c+1], grid[r][c], grid[r+1][c-1]}

			// Center must be 'A' and both diagonals must be valid permutations of "MAS"
			return grid[r][c] == 'A' &&
				isDiagonalXMas(topLeftToBottomRight) &&
				isDiagonalXMas(topRightToBottomLeft)
		}
		return false
	}

	// Iterate over the grid to find all "X-MAS" patterns
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if checkXMas(r, c) {
				results = append(results, [2]int{r, c})
			}
		}
	}

	return results
}
