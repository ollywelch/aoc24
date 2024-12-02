package day2

import (
	"math"
)

func IsSafe(nums []int) bool {
	var isDecreasing bool
	for i := range nums {
		if i == 0 {
			isDecreasing = nums[0] >= nums[1]
			continue
		}
		if isDecreasing {
			if nums[i] >= nums[i-1] {
				return false
			}
		} else {
			if nums[i] <= nums[i-1] {
				return false
			}
		}
		if math.Abs(float64(nums[i]-nums[i-1])) > 3 {
			return false
		}
	}
	return true
}

func IsSafeWithDampener(nums []int) bool {
	// being lazy, I think we can do this a bit better :)
	if IsSafe(nums) {
		return true
	}
	for i := range nums {
		numsWithRemoved := removeIndex(nums, i)
		if IsSafe(numsWithRemoved) {
			return true
		}
	}
	return false
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
