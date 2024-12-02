package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ollywelch/aoc24/day5"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	before := []int{}
	after := []int{}
	seqs := [][]int{}
	for _, line := range strings.Split(string(inputBytes), "\n") {
		if strings.Contains(line, "|") {
			splitLine := strings.Split(line, "|")
			if len(splitLine) != 2 {
				log.Fatal(fmt.Errorf("line %s had split length %d", line, len(splitLine)))
			}
			b, err := strconv.Atoi(splitLine[0])
			if err != nil {
				log.Fatal(err)
			}
			a, err := strconv.Atoi(splitLine[1])
			if err != nil {
				log.Fatal(err)
			}
			before = append(before, b)
			after = append(after, a)
			continue
		}
		if strings.Contains(line, ",") {
			seq := []int{}
			splitLine := strings.Split(line, ",")
			for _, item := range splitLine {
				num, err := strconv.Atoi(item)
				if err != nil {
					log.Fatal(err)
				}
				seq = append(seq, num)
			}
			seqs = append(seqs, seq)
		}
	}

	count := 0
	fixedCount := 0
	for _, seq := range seqs {
		if day5.IsValidSequence(seq, before, after) {
			count += day5.MiddleNumber(seq)
		} else {
			fixedCount += day5.MiddleNumber(day5.FixInvalidSequence(seq, before, after))
		}
	}
	fmt.Println("Total for valid sequences:", count)
	fmt.Println("Total for fixed invalid sequences:", fixedCount)
}
