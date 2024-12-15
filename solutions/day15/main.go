package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ollywelch/aoc24/day15"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day15.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid, moves := day15.ParseInput(string(inputBytes))

	grid = day15.DoMoves(grid, moves)

	fmt.Println("Part 1:", day15.CalculateGPS(grid))
}
