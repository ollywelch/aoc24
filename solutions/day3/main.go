package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ollywelch/aoc24/day3"
)

func main() {
	inputBytes, err := os.ReadFile("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	nums1, nums2 := day3.Parse(string(inputBytes))
	total := 0
	if len(nums1) != len(nums2) {
		log.Fatal(fmt.Errorf("nums1 length did not match nums2: %d != %d", len(nums1), len(nums2)))
	}
	for i := range nums1 {
		total += nums1[i] * nums2[i]
	}
	fmt.Println("Total:", total)
}
