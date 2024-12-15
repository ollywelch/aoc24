package day15

import (
	"strings"
)

const (
	MoveUp = iota
	MoveDown
	MoveLeft
	MoveRight
)

type Move int

type Coordinate struct {
	Row    int
	Column int
}

func ParseInput(input string) ([]string, []Move) {
	grid := []string{}
	moves := []Move{}

	splitInput := strings.Split(input, "\n")
	for _, line := range splitInput {
		if strings.HasPrefix(line, "#") {
			grid = append(grid, line)
			continue
		}
		if line == "" {
			continue
		}
		for _, char := range line {
			switch char {
			case '^':
				moves = append(moves, MoveUp)
			case '>':
				moves = append(moves, MoveRight)
			case '<':
				moves = append(moves, MoveLeft)
			case 'v':
				moves = append(moves, MoveDown)
			}
		}
	}
	return grid, moves
}

func CalculateGPS(grid []string) int {
	total := 0
	for i := range grid {
		for j := range grid[i] {
			if rune(grid[i][j]) == 'O' {
				total += 100*i + j
			}
		}
	}
	return total
}

func DoMoves(grid []string, moves []Move) []string {
	var position Coordinate
	for i := range grid {
		for j := range grid[i] {
			if rune(grid[i][j]) == '@' {
				position.Row = i
				position.Column = j
			}
		}
	}

	for _, move := range moves {
		grid, position = Next(grid, position, move)
	}

	return grid
}

func Next(grid []string, position Coordinate, move Move) ([]string, Coordinate) {
	coordFunc := coordinateFunc(move)
	movePos := coordFunc(position)
	nextPos := movePos
	var finalChar rune
	for {
		nextChar := rune(grid[nextPos.Row][nextPos.Column])

		if nextChar == '.' || nextChar == '#' {
			finalChar = nextChar
			break
		}

		nextPos = coordFunc(nextPos)
	}
	if nextPos == movePos {
		if finalChar == '#' {
			return grid, position
		}
		grid = setCoordinate(grid, position.Row, position.Column, '.')
		grid = setCoordinate(grid, movePos.Row, movePos.Column, '@')
		return grid, movePos
	} else {
		if finalChar == '.' {
			grid = setCoordinate(grid, position.Row, position.Column, '.')
			grid = setCoordinate(grid, movePos.Row, movePos.Column, '@')
			grid = setCoordinate(grid, nextPos.Row, nextPos.Column, 'O')
			return grid, movePos
		} else {
			return grid, position
		}
	}
}

func coordinateFunc(move Move) func(Coordinate) Coordinate {
	var coordFunc func(Coordinate) Coordinate
	switch move {
	case MoveUp:
		coordFunc = func(c Coordinate) Coordinate {
			c.Row--
			return c
		}
	case MoveDown:
		coordFunc = func(c Coordinate) Coordinate {
			c.Row++
			return c
		}
	case MoveLeft:
		coordFunc = func(c Coordinate) Coordinate {
			c.Column--
			return c
		}
	case MoveRight:
		coordFunc = func(c Coordinate) Coordinate {
			c.Column++
			return c
		}
	}
	return coordFunc
}

func setCoordinate(grid []string, i, j int, r rune) []string {
	row := grid[i]
	row = replaceAtIndex(row, r, j)
	grid[i] = row
	return grid
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
