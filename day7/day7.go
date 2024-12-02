package day7

import (
	"fmt"
	"strconv"
)

func IsValidWithConcatenation(target int, inputs []int) bool {
	if len(inputs) < 1 {
		return false
	}
	if (inputs[0]) > target {
		return false
	}
	if len(inputs) == 1 {
		return inputs[0] == target
	}
	addInputs := removeIndex(inputs, 0)
	addInputs[0] = addInputs[0] + inputs[0]
	mulInputs := removeIndex(inputs, 0)
	mulInputs[0] = mulInputs[0] * inputs[0]
	testConcat := true
	concatInputs := removeIndex(inputs, 0)
	var err error
	concatInputs[0], err = strconv.Atoi(fmt.Sprint(inputs[0]) + fmt.Sprint(concatInputs[0]))
	if err != nil {
		testConcat = false
	}
	return IsValidWithConcatenation(target, addInputs) ||
		IsValidWithConcatenation(target, mulInputs) ||
		(testConcat && IsValidWithConcatenation(target, concatInputs))
}

func IsValid(target int, inputs []int) bool {
	if len(inputs) < 1 {
		return false
	}
	if (inputs[0]) > target {
		return false
	}
	if len(inputs) == 1 {
		return inputs[0] == target
	}
	addInputs := removeIndex(inputs, 0)
	addInputs[0] = addInputs[0] + inputs[0]
	mulInputs := removeIndex(inputs, 0)
	mulInputs[0] = mulInputs[0] * inputs[0]
	return IsValid(target, addInputs) ||
		IsValid(target, mulInputs)
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
