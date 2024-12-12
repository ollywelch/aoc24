package day11

import (
	"strconv"
)

// cache per depth level of root nodes to their counts
var depthCache = map[int]map[int]int{}

func CountStones(in []int, repeats int) int {
	count := 0
	for _, num := range in {
		count += CountWithDepth(num, repeats)
	}
	return count
}

func CountWithDepth(root int, depth int) int {
	if depth == 0 {
		return 1
	}
	cache, ok := depthCache[depth]
	if !ok {
		cache = make(map[int]int)
	}
	if cachedResult, ok := cache[root]; ok {
		return cachedResult
	}
	children := DoBlink(root)
	count := CountWithDepth(children[0], depth-1)
	if len(children) > 1 {
		count += CountWithDepth(children[1], depth-1)
	}
	cache[root] = count
	depthCache[depth] = cache
	return count
}

func DoBlink(num int) []int {
	if num == 0 {
		return []int{1}
	}

	strNum := strconv.Itoa(num)
	nDigits := len(strNum)
	if nDigits%2 == 1 {
		return []int{num * 2024}
	}

	splitAt := nDigits / 2
	leftDigits := strNum[:splitAt]
	rightDigits := strNum[splitAt:]
	left, err := strconv.Atoi(leftDigits)
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(rightDigits)
	if err != nil {
		panic(err)
	}
	return []int{left, right}
}
