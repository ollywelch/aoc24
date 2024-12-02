package day5

import (
	"math"
	"slices"
)

func FixInvalidSequence(seq []int, before []int, after []int) []int {
	slices.SortFunc(seq, func(a, b int) int {
		for i := range before {
			if before[i] == a && after[i] == b {
				return -1
			}
		}
		return 0
	})
	return seq
}

func IsValidSequence(seq []int, before []int, after []int) bool {
	for i := range before {
		afterIsFirst := false
		for _, num := range seq {
			if num == before[i] {
				if afterIsFirst {
					return false
				}
				break
			}
			if num == after[i] {
				afterIsFirst = true
			}
		}
	}
	return true
}

func MiddleNumber(seq []int) int {
	idx := int(math.Floor(float64(len(seq)+1)/2)) - 1
	return seq[idx]
}
