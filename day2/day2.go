package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	reports := createArrayFromFile()
	fmt.Printf("Reports: %v\n", reports)
}

func createArrayFromFile() [][]int {
	filePath := "day2/input.txt"
	var reports [][]int

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // Ensure the file is closed when the function exits

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the text of the current line
		levelStr := strings.Split(line, " ")
		var levels []int
		for i := 0; i < len(levelStr); i++ {
			level, err := strconv.Atoi(levelStr[i])
			if err != nil {
				panic(err)
			}
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}
	return reports
}
