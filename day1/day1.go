package day1

import (
	"math"
	"slices"
)

func CalculateDistance(nums1, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	var distance float64
	if len(nums1) != len(nums2) {
		return -1
	}
	for i := range nums1 {
		distance += math.Abs(float64(nums1[i] - nums2[i]))
	}
	return int(distance)
}

func CalculateSimilarity(nums1, nums2 []int) int {
	if len(nums1) != len(nums2) {
		return -1
	}
	var totalSimilarity int
	for _, num1 := range nums1 {
		// Check how many times it appears in nums2
		count := 0
		for _, num2 := range nums2 {
			if num1 == num2 {
				count++
			}
		}

		totalSimilarity += num1 * count
	}
	return totalSimilarity
}
