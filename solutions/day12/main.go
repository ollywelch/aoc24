package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ollywelch/aoc24/day12"
)

func main() {
	input, err := os.Open("inputs/day12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	grid := []string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", day12.CalculatePrice(grid))
	fmt.Println("Part 2:", day12.CalculatePrice2(grid))
}
