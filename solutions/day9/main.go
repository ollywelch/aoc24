package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ollywelch/aoc24/day9"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Checksum:", day9.CalculateChecksum(string(inputBytes)))
	fmt.Println("Checksum new:", day9.CalculateChecksumNew(string(inputBytes)))
}
