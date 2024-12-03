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
	nums3, nums4 := day3.Parse2(string(inputBytes))
	total := 0
	total2 := 0
	if len(nums1) != len(nums2) {
		log.Fatal(fmt.Errorf("nums1 length did not match nums2: %d != %d", len(nums1), len(nums2)))
	}
	if len(nums3) != len(nums4) {
		log.Fatal(fmt.Errorf("nums3 length did not match nums4: %d != %d", len(nums3), len(nums4)))
	}
	for i := range nums1 {
		total += nums1[i] * nums2[i]
	}
	for i := range nums3 {
		total2 += nums3[i] * nums4[i]
	}
	fmt.Println("Total:", total)
	fmt.Println("Total 2:", total2)
}
