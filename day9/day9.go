package day9

import (
	"strconv"
)

func CalculateChecksumNew(input string) int {
	checkSum := 0
	sparse := toSparse(input)
	right := len(sparse) - 1
	for {
		if right < 0 {
			break
		}
		if sparse[right] < 0 {
			right--
			continue
		}
		rightCount := 0
		leftCount := 0
		left := 0
		innerRight := right
		for {
			// keep going until the one on the right has a different value
			if innerRight < 0 || sparse[innerRight] != sparse[right] {
				break
			}
			rightCount++
			innerRight--
		}
		// now go left to right and look for a gap of rightCount spaces
		for {
			if sparse[left] < 0 {
				leftCount++
			} else {
				leftCount = 0
			}
			if leftCount >= rightCount || left >= right {
				break
			}
			left++
		}
		if leftCount >= rightCount {
			for i := 0; i < rightCount; i++ {
				sparse[left-i] = sparse[right-i]
				sparse[right-i] = -1
			}
		}
		right -= rightCount
	}
	for i, num := range sparse {
		if num < 0 {
			continue
		}
		checkSum += i * num
	}
	return checkSum
}

func CalculateChecksum(input string) int {
	checkSum := 0
	sparse := toSparse(input)
	left := 0
	right := len(sparse) - 1
	for {
		if left > right {
			break
		}
		if sparse[left] >= 0 {
			checkSum += left * sparse[left]
			left++
			continue
		}
		if sparse[right] < 0 {
			right--
			continue
		}
		checkSum += left * sparse[right]
		left++
		right--
	}

	return checkSum
}

func toSparse(input string) []int {
	sparse := []int{}
	for i := range input {
		repeats, _ := strconv.Atoi(string(input[i]))
		toAdd := -1
		if i%2 == 0 {
			toAdd = i / 2
		}
		for range repeats {
			sparse = append(sparse, toAdd)
		}
	}
	return sparse
}
