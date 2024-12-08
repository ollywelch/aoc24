package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ollywelch/aoc24/day8"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(inputBytes), "\n")
	antiNodes := day8.FindAntiNodes(input)
	allAntiNodes := day8.FindAllAntiNodes(input)
	fmt.Println("Total anti nodes:", len(antiNodes))
	fmt.Println("Total anti nodes with resonant harmonics", len(allAntiNodes))
}
