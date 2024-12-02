package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ollywelch/aoc24/day10"
)

func main() {
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(day10.CalculateScore(grid))
	fmt.Println(day10.CalculateRating(grid))
}
