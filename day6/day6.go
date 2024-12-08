package day6

import (
	"advent_of_code_2024/helper"
	"fmt"
)

type GameStates struct {
	Direction     rune
	OutOfBounds   bool
	TilesVisisted map[struct{ x, y int }]struct{}
	GuardPos      struct {
		x, y int
	}
}

// ^ = guard . = empty space # = obstruction X = path
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

// Day6 solves the puzzles for day 6
// It finds the number of tiles visited by the guard
// and the number of infinite loops the guard gets stuck in
func Day6() {
	rawInput := helper.MatrixFromFile("day6/input.txt")
	x, y := findPlayerPosition(rawInput, gameTiles.Guard)
	if x == -1 || y == -1 {
		panic("Guard not found")
	}

	gameState := setUpGameStates(x, y)
	startGame(gameState, rawInput)
	fmt.Printf("Visisted: %d\n", len(gameState.TilesVisisted))

	getInfinteLoop := getGuardStuckInLoop(rawInput, gameState)
	fmt.Printf("Infinite loops: %d\n", getInfinteLoop)
}

// startGame starts the game and returns true if the guard goes out of bounds
// or false if the guard completes the game without going out of bounds
func startGame(gameState GameStates, rawInput [][]rune) bool {
	maxMoves := len(rawInput) * len(rawInput[0])
	count := 0
	for !gameState.OutOfBounds {
		if count == maxMoves {
			return false
		}
		nextX, nextY := move(gameState.GuardPos.x, gameState.GuardPos.y, gameState)
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

// getGuardStuckInLoop finds the number of infinite loops the guard gets stuck in
func getGuardStuckInLoop(grid [][]rune, firstRun GameStates) int {
	infiniteLoop := 0

	for pos := range firstRun.TilesVisisted {
		copyGrid := make([][]rune, len(grid))
		for i := range grid {
			copyGrid[i] = make([]rune, len(grid[i]))
			copy(copyGrid[i], grid[i])
		}
		x, y := pos.x, pos.y
		copyGrid[x][y] = gameTiles.Obstruction
		gameState := setUpGameStates(firstRun.GuardPos.x, firstRun.GuardPos.y)
		if !startGame(gameState, copyGrid) {
			infiniteLoop++
		}
	}
	return infiniteLoop
}

// setUpGameStates initializes the game states
func setUpGameStates(startingPositionX int, startingPositionY int) GameStates {
	return GameStates{
		Direction:     'N',
		OutOfBounds:   false,
		TilesVisisted: make(map[struct{ x, y int }]struct{}),
		GuardPos:      struct{ x, y int }{startingPositionX, startingPositionY},
	}
}

// turn90Degrees turns the guard 90 degrees to the left or right based on the current direction
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

// move moves the guard in the specified direction
func move(x, y int, gameStates GameStates) (int, int) {
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

// findPlayerPosition finds the player's position in the grid
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

// isMoveOutOfBounds checks if the next move is out of bounds
func isMoveOutOfBounds(grid [][]rune, x, y int) bool {
	return x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0])
}
