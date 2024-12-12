package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ollywelch/aoc24/day11"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	strNums := strings.Split(string(inputBytes), " ")
	input := make([]int, len(strNums))
	for i := range input {
		var err error
		input[i], err = strconv.Atoi(strNums[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Number of stones after 75:", day11.CountStones(input, 75))
}
