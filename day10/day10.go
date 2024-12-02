package day10

import (
	"strconv"
)

type Coordinates map[int]map[int]struct{}
type CoordinateScores map[int]map[int]int

func CalculateScore(grid []string) int {
	// mapping of numbers to coordinates they appear at
	// get initialised with a score of 1 (i.e. 9's cause an increase in score of 1)
	coordinateMap := map[int]Coordinates{}

	// initialise coordinates from grid
	for i := range grid {
		for j := range grid[i] {
			num, err := strconv.Atoi(string(grid[i][j]))
			if err != nil {
				panic("non integer cell encountered!")
			}
			numCoordinates, ok := coordinateMap[num]
			if !ok {
				coordinateMap[num] = Coordinates{i: map[int]struct{}{j: {}}}
				continue
			}
			row, ok := numCoordinates[i]
			if !ok {
				numCoordinates[i] = map[int]struct{}{j: {}}
				continue
			}
			row[j] = struct{}{}
		}
	}
	count := 0
	for row := range coordinateMap[0] {
		for column := range coordinateMap[0][row] {
			trailMap := map[int]Coordinates{}
			trailMap[0] = Coordinates{row: map[int]struct{}{column: {}}}
			for i := 1; i <= 9; i++ {
				// add the reachable i's to the trailMap
				iCoordinates := coordinateMap[i]
				trailCoordinates := trailMap[i-1]
				nextCoordinates := make(Coordinates)
				for row := range trailCoordinates {
					for column := range trailCoordinates[row] {
						if exists(row, column+1, iCoordinates) {
							nextCoordinates = insert(row, column+1, nextCoordinates)
						}
						if exists(row, column-1, iCoordinates) {
							nextCoordinates = insert(row, column-1, nextCoordinates)
						}
						if exists(row+1, column, iCoordinates) {
							nextCoordinates = insert(row+1, column, nextCoordinates)
						}
						if exists(row-1, column, iCoordinates) {
							nextCoordinates = insert(row-1, column, nextCoordinates)
						}
					}
				}
				trailMap[i] = nextCoordinates
			}
			for _, row := range trailMap[9] {
				for range row {
					count++
				}
			}
		}
	}

	return count
}

func exists(i, j int, coordinates Coordinates) bool {
	if row, ok := coordinates[i]; ok {
		if _, ok := row[j]; ok {
			return true
		}
	}
	return false
}

func insert(i, j int, coordinates Coordinates) Coordinates {
	if row, ok := coordinates[i]; ok {
		row[j] = struct{}{}
	} else {
		coordinates[i] = map[int]struct{}{j: {}}
	}
	return coordinates
}

func CalculateRating(grid []string) int {
	// mapping of numbers to coordinates they appear at
	// get initialised with a score of 1 (i.e. 9's cause an increase in score of 1)
	coordinates := map[int]CoordinateScores{}

	// initialise coordinates from grid
	for i := range grid {
		for j := range grid[i] {
			num, err := strconv.Atoi(string(grid[i][j]))
			if err != nil {
				panic("non integer cell encountered!")
			}
			score := 0
			if num == 9 {
				score = 1
			}
			numCoordinates, ok := coordinates[num]
			if !ok {
				coordinates[num] = CoordinateScores{i: map[int]int{j: score}}
				continue
			}
			row, ok := numCoordinates[i]
			if !ok {
				numCoordinates[i] = map[int]int{j: score}
				continue
			}
			row[j] = score
		}
	}

	// now go from 8->0 calculating the scores based on adjacent cells
	// new score = sum(adjacent n+1's * their score)
	for i := 8; i >= 0; i-- {
		nCoordinates := coordinates[i]
		nPlusOneCoordinates := coordinates[i+1]
		for row, columns := range nCoordinates {
			for column := range columns {
				// check same row first
				if sameRow, ok := nPlusOneCoordinates[row]; ok {
					if score, ok := sameRow[column+1]; ok {
						coordinates[i][row][column] += score
					}
					if score, ok := sameRow[column-1]; ok {
						coordinates[i][row][column] += score
					}
				}
				if belowRow, ok := nPlusOneCoordinates[row+1]; ok {
					if score, ok := belowRow[column]; ok {
						coordinates[i][row][column] += score
					}
				}
				if aboveRow, ok := nPlusOneCoordinates[row-1]; ok {
					if score, ok := aboveRow[column]; ok {
						coordinates[i][row][column] += score
					}
				}
			}
		}
	}

	total := 0
	trailheads := coordinates[0]
	for _, row := range trailheads {
		for _, score := range row {
			total += score
		}
	}
	return total
}
