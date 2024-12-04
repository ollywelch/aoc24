package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ollywelch/aoc24/day4"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := strings.Split(string(inputBytes), "\n")
	fmt.Println("XMAS Count:", day4.CountWords(grid, "XMAS"))
	fmt.Println("X-MAS Count:", day4.CountXMAS(grid))
}
