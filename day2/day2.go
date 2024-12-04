package day2

import (
	"advent_of_code_2024/helper"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	reports := createArrayFromFile()
	getSafeReports(reports)
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

func getSafeReports(reports [][]int) int {
	badReports := 0
	goodReports := 0
	for i := 0; i < len(reports); i++ {
		// var last_number int
		var is_incrementing bool
		if reports[i][0] > reports[i][1] {
			is_incrementing = false
		} else {
			is_incrementing = true
		}
		for j := 0; j < len(reports[i]); j++ {
			if j != len(reports[i])-1 {
				next_num := reports[i][j+1]
				if !checkIfIncrementing(reports[i][j], next_num, is_incrementing) {
					badReports++
					break
				}
				if !checkIfLevelsDiffer(reports[i][j], next_num) {
					badReports++
					break
				} else if j == len(reports[i])-2 && checkIfLevelsDiffer(reports[i][j], next_num) {
					goodReports++

				}
			}
		}
	}
	fmt.Printf("Bad Reports: %d\n", badReports)
	fmt.Printf("Good Reports: %d\n", goodReports)
	return goodReports
}

func checkIfLevelsDiffer(currentValue int, next_value int) bool {
	if currentValue == next_value {
		return false
	}
	difference := helper.Abs(currentValue - next_value)
	return difference <= 3
}

func checkIfIncrementing(current_value int, next_value int, is_incrementing bool) bool {
	if is_incrementing {
		if current_value > next_value {
			return false
		} else {
			return true
		}
	} else {
		if current_value < next_value {
			return false
		} else {
			return true
		}
	}
}
