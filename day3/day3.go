package day3

import (
	"strconv"
	"unicode"
)

func Parse2(input string) ([]int, []int) {
	nums1 := []int{}
	nums2 := []int{}

	// Only have to loop up to len(input) - 8
	// as mul(1,1) is 8 characters long
	mulEnabled := true
	for i := range input[:len(input)-8] {
		if input[i:i+4] == `do()` {
			mulEnabled = true
			continue
		}
		if input[i:i+7] == `don't()` {
			mulEnabled = false
			continue
		}
		// First look for `mul(`
		if mulEnabled && input[i:i+4] == `mul(` {
			// now loop over the next characters while we keep finding digits
			// break the loop if we find a non-digit/comma
			// parse the result if we found a comma
			digitStack1 := ``
			digitStack2 := ``
			for _, char := range input[i+4:] {
				if !unicode.IsDigit(char) {
					break
				}
				digitStack1 += string(char)
			}
			if len(digitStack1) == 0 || len(digitStack1) > 3 {
				continue
			}
			breakChar1 := input[i+4+len(digitStack1)]
			if rune(breakChar1) != ',' {
				continue
			}
			if i+4+len(digitStack1) > len(input)-1 {
				break
			}
			for _, char := range input[i+4+len(digitStack1)+1:] {
				if !unicode.IsDigit(char) {
					break
				}
				digitStack2 += string(char)
			}
			if len(digitStack2) == 0 || len(digitStack2) > 3 {
				continue
			}
			breakChar2 := input[i+4+len(digitStack1)+1+len(digitStack2)]
			if rune(breakChar2) != ')' {
				continue
			}
			num1, err := strconv.Atoi(digitStack1)
			if err != nil {
				panic(err)
			}
			num2, err := strconv.Atoi(digitStack2)
			if err != nil {
				panic(err)
			}
			nums1 = append(nums1, num1)
			nums2 = append(nums2, num2)
		}
	}
	return nums1, nums2
}

func Parse(input string) ([]int, []int) {
	nums1 := []int{}
	nums2 := []int{}

	// Only have to loop up to len(input) - 8
	// as mul(1,1) is 8 characters long
	for i := range input[:len(input)-8] {
		// First look for `mul(`
		if input[i:i+4] == `mul(` {
			// now loop over the next characters while we keep finding digits
			// break the loop if we find a non-digit/comma
			// parse the result if we found a comma
			digitStack1 := ``
			digitStack2 := ``
			for _, char := range input[i+4:] {
				if !unicode.IsDigit(char) {
					break
				}
				digitStack1 += string(char)
			}
			if len(digitStack1) == 0 || len(digitStack1) > 3 {
				continue
			}
			breakChar1 := input[i+4+len(digitStack1)]
			if rune(breakChar1) != ',' {
				continue
			}
			if i+4+len(digitStack1) > len(input)-1 {
				break
			}
			for _, char := range input[i+4+len(digitStack1)+1:] {
				if !unicode.IsDigit(char) {
					break
				}
				digitStack2 += string(char)
			}
			if len(digitStack2) == 0 || len(digitStack2) > 3 {
				continue
			}
			breakChar2 := input[i+4+len(digitStack1)+1+len(digitStack2)]
			if rune(breakChar2) != ')' {
				continue
			}
			num1, err := strconv.Atoi(digitStack1)
			if err != nil {
				panic(err)
			}
			num2, err := strconv.Atoi(digitStack2)
			if err != nil {
				panic(err)
			}
			nums1 = append(nums1, num1)
			nums2 = append(nums2, num2)
		}
	}
	return nums1, nums2
}
