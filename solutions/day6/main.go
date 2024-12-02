package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ollywelch/aoc24/day6"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputGrid := strings.Split(string(inputBytes), "\n")
	grid := make([]string, 0)
	grid = append(grid, inputGrid...)
	day6.TrackGuard(grid)
	posCount := 0
	for i := range grid {
		for j := range grid[i] {
			if rune(grid[i][j]) == 'X' {
				posCount++
			}
		}
	}
	fmt.Println("Total guard positions:", posCount)
	// v brute force method, should probably find a quicker way
	// look for patterns in how we can force a loop rather than
	// just try every visited cell. This is slow
	potentialObstacleCount := 0
	for i := range grid {
		for j := range grid[i] {
			if rune(grid[i][j]) == 'X' && rune(inputGrid[i][j]) != '^' {
				copiedGrid := make([]string, 0)
				copiedGrid = append(copiedGrid, inputGrid...)
				copiedGrid[i] = day6.ReplaceAtIndex(copiedGrid[i], '#', j)
				copiedGrid, hasLoop := day6.TrackGuard(copiedGrid)
				if hasLoop {
					potentialObstacleCount++
				}
			}
		}
	}
	fmt.Println("Potential obstacle count:", potentialObstacleCount)
}
