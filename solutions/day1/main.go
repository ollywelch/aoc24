package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ollywelch/aoc24/day1"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputLines := strings.Split(string(inputBytes), "\n")
	nums1 := []int{}
	nums2 := []int{}
	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		num1, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(splitLine[len(splitLine)-1])
		if err != nil {
			log.Fatal(err)
		}
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}

	fmt.Println("Distance:", day1.CalculateDistance(nums1, nums2))
	fmt.Println("Similarity:", day1.CalculateSimilarity(nums1, nums2))
}
