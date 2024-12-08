package day6

import (
	"advent_of_code_2024/helper"
	"fmt"
)

// ^ = guard . = empty space # = obstruction X = path
type GameStates struct {
	Guard         rune
	EmptySpace    rune
	Obstruction   rune
	Path          rune
	Direction     rune
	OutOfBounds   bool
	TilesVisisted map[struct{ x, y int }]struct{}
	GuardPos      struct {
		x, y int
	}
}

// Day6 simulates a game where a guard navigates through a grid based on specific rules.
// The guard starts at a position marked by '^' and moves in a direction ('N' for North).
// The grid contains empty spaces ('.'), obstructions ('#'), and paths ('X').
// The guard moves until it goes out of bounds or encounters an obstruction, at which point it turns 90 degrees to the right.
// The function reads the initial grid from "day6/input.txt", finds the guard's starting position,
// and then processes the guard's movements, marking the path taken and counting the number of unique tiles visited.
// Finally, it prints the resulting grid and the number of unique tiles visited.
func Day6() {
	gameStates := GameStates{
		Guard:         '^',
		EmptySpace:    '.',
		Obstruction:   '#',
		Path:          'X',
		Direction:     'N',
		OutOfBounds:   false,
		TilesVisisted: make(map[struct{ x, y int }]struct{}),
	}
	rawInput := helper.MatrixFromFile("day6/input.txt")
	x, y := findPlayerPosition(rawInput, gameStates.Guard)
	if x == -1 || y == -1 {
		panic("Guard not found")
	}
	gameStates.GuardPos.x = x
	gameStates.GuardPos.y = y
	for !gameStates.OutOfBounds {
		nextX, nextY := move(rawInput, gameStates.GuardPos.x, gameStates.GuardPos.y, gameStates)
		if isMoveOutOfBounds(rawInput, nextX, nextY) {
			gameStates.OutOfBounds = true
			break
		} else if rawInput[nextX][nextY] == gameStates.Obstruction {
			gameStates.Direction = turn90Degrees(gameStates.Direction, 'R')
		} else {
			gameStates.GuardPos.x = nextX
			gameStates.GuardPos.y = nextY
			rawInput[nextX][nextY] = gameStates.Path
			if _, visited := gameStates.TilesVisisted[struct{ x, y int }{nextX, nextY}]; !visited {
				gameStates.TilesVisisted[struct{ x, y int }{nextX, nextY}] = struct{}{}
			}
		}
	}
	printMap(rawInput)
	fmt.Printf("Visisted: %d\n", len(gameStates.TilesVisisted))
}

// turn90Degrees
func turn90Degrees(direction rune, turn rune) rune {
	if turn == 'L' {
		switch direction {
		case 'N':
			return 'W'
		case 'S':
			return 'E'
		case 'E':
			return 'N'
		case 'W':
			return 'S'
		}
	} else if turn == 'R' {
		switch direction {
		case 'N':
			return 'E'
		case 'S':
			return 'W'
		case 'E':
			return 'S'
		case 'W':
			return 'N'
		}
	}
	return ' '
}

func move(grid [][]rune, x, y int, gameStates GameStates) (int, int) {
	switch gameStates.Direction {
	case 'N':
		x--
	case 'S':
		x++
	case 'E':
		y++
	case 'W':
		y--
	}
	return x, y
}

func findPlayerPosition(grid [][]rune, player rune) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == player {
				return i, j
			}
		}
	}
	return -1, -1
}

func isMoveOutOfBounds(grid [][]rune, x, y int) bool {
	return x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0])
}

func printMap(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			print(string(cell))
		}
		println()
	}
}
