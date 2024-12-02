package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ollywelch/aoc24/day2"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputLines := strings.Split(string(inputBytes), "\n")
	safeCount := 0
	dampenerSafeCount := 0
	for _, line := range inputLines {
		strNums := strings.Split(line, " ")
		nums := []int{}
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		if day2.IsSafe(nums) {
			safeCount++
		}
		if day2.IsSafeWithDampener(nums) {
			dampenerSafeCount++
		}
	}
	fmt.Println("Safe count:", safeCount)
	fmt.Println("Safe count with dampener:", dampenerSafeCount)
}
