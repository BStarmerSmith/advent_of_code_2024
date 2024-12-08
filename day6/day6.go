package day6

import (
	"advent_of_code_2024/helper"
	"fmt"
)

// ^ = guard . = empty space # = obstruction X = path
type GameStates struct {
	Direction     rune
	OutOfBounds   bool
	TilesVisisted map[struct{ x, y int }]struct{}
	GuardPos      struct {
		x, y int
	}
}

type GameTiles struct {
	Guard       rune
	EmptySpace  rune
	Obstruction rune
	Path        rune
}

var gameTiles GameTiles

func init() {
	gameTiles = GameTiles{
		Guard:       '^',
		EmptySpace:  '.',
		Obstruction: '#',
		Path:        'X',
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
	rawInput := helper.MatrixFromFile("day6/input.txt")
	x, y := findPlayerPosition(rawInput, gameTiles.Guard)
	if x == -1 || y == -1 {
		panic("Guard not found")
	}

	gameState := setUpGameStates(rawInput, x, y)
	startGame(gameState, rawInput)
	fmt.Printf("Visisted: %d\n", len(gameState.TilesVisisted))

	getInfinteLoop := getGuardStuckInLoop(rawInput, x, y)
	fmt.Printf("Infinite loops: %d\n", getInfinteLoop)
}

func startGame(gameState GameStates, rawInput [][]rune) bool {
	maxMoves := len(rawInput) * len(rawInput[0])
	count := 0
	for !gameState.OutOfBounds {
		if count == maxMoves {
			return false
		}
		nextX, nextY := move(rawInput, gameState.GuardPos.x, gameState.GuardPos.y, gameState)
		if isMoveOutOfBounds(rawInput, nextX, nextY) {
			gameState.OutOfBounds = true
			break
		} else if rawInput[nextX][nextY] == gameTiles.Obstruction {
			gameState.Direction = turn90Degrees(gameState.Direction, 'R')
		} else {
			gameState.GuardPos.x = nextX
			gameState.GuardPos.y = nextY
			rawInput[nextX][nextY] = gameTiles.Path
			if _, visited := gameState.TilesVisisted[struct{ x, y int }{nextX, nextY}]; !visited {
				gameState.TilesVisisted[struct{ x, y int }{nextX, nextY}] = struct{}{}
			}
		}
		count++
	}
	return true
}

func getGuardStuckInLoop(grid [][]rune, startingX int, startingY int) int {
	infiniteLoop := 0
	for i, row := range grid {
		for j, cell := range row {
			totalCells := len(grid) * len(grid[0])
			processedCells := i*len(row) + j + 1
			if processedCells*10%totalCells == 0 {
				fmt.Printf("Processed %d%% of cells\n", processedCells*100/totalCells)
			}
			if cell == '.' {
				oldCell := grid[i][j]
				grid[i][j] = gameTiles.Obstruction
				gameState := setUpGameStates(grid, startingX, startingY)
				if !startGame(gameState, grid) {
					infiniteLoop++
				}
				grid[i][j] = oldCell
			}
		}
	}
	return infiniteLoop
}

func setUpGameStates(grid [][]rune, startingPositionX int, startingPositionY int) GameStates {
	return GameStates{
		Direction:     'N',
		OutOfBounds:   false,
		TilesVisisted: make(map[struct{ x, y int }]struct{}),
		GuardPos:      struct{ x, y int }{startingPositionX, startingPositionY},
	}
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
