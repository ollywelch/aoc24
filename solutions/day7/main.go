package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ollywelch/aoc24/day7"
)

func main() {
	input, err := os.Open("inputs/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	total := 0
	totalWithConcat := 0

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		before, after, _ := strings.Cut(line, ": ")
		target, err := strconv.Atoi(before)
		if err != nil {
			log.Fatal(err)
		}
		strInputs := strings.Split(after, " ")
		inputs := []int{}
		for _, strInput := range strInputs {
			input, err := strconv.Atoi(strInput)
			if err != nil {
				log.Fatal(err)
			}
			inputs = append(inputs, input)
		}
		if day7.IsValid(target, inputs) {
			total += target
			totalWithConcat += target
		} else {
			if day7.IsValidWithConcatenation(target, inputs) {
				totalWithConcat += target
			}
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total:", total)
	fmt.Println("Total with concatenation:", totalWithConcat)
}
