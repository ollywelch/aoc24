package day7

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value    int
	Children []Node
}

func IsValidWithConcatenation(target int, inputs []int) bool {
	if len(inputs) < 1 {
		return false
	}
	if len(inputs) == 1 {
		return inputs[0] == target
	}
	tree := &Node{
		Value: inputs[0],
		Children: []Node{
			*subTree(target, inputs[1:], inputs[0], '+'),
			*subTree(target, inputs[1:], inputs[0], '*'),
			*subTree(target, inputs[1:], inputs[0], '|'),
		},
	}
	leaves := leafNodes(tree)
	for _, leaf := range leaves {
		if leaf == target {
			return true
		}
	}
	return false
}

func IsValid(target int, inputs []int) bool {
	if len(inputs) < 1 {
		return false
	}
	if len(inputs) == 1 {
		return inputs[0] == target
	}
	tree := &Node{
		Value: inputs[0],
		Children: []Node{
			*subTree(target, inputs[1:], inputs[0], '+'),
			*subTree(target, inputs[1:], inputs[0], '*'),
		},
	}
	leaves := leafNodes(tree)
	for _, leaf := range leaves {
		if leaf == target {
			return true
		}
	}
	return false
}

func subTree(target int, inputs []int, value int, op rune) *Node {
	switch op {
	case '+':
		value += inputs[0]
	case '*':
		value *= inputs[0]
	case '|':
		concatValue, err := strconv.Atoi(fmt.Sprint(value) + fmt.Sprint(inputs[0]))
		if err != nil {
			return nil
		}
		value = concatValue
	default:
		return nil
	}
	n := &Node{
		Value: value,
	}
	if len(inputs) < 2 || value >= target {
		return n
	}
	if s := subTree(target, inputs[1:], value, '+'); s != nil {
		n.Children = append(n.Children, *s)
	}
	if s := subTree(target, inputs[1:], value, '*'); s != nil {
		n.Children = append(n.Children, *s)
	}
	if s := subTree(target, inputs[1:], value, '|'); s != nil {
		n.Children = append(n.Children, *s)
	}
	return n
}

func leafNodes(tree *Node) []int {
	leaves := []int{}
	if len(tree.Children) == 0 {
		leaves = append(leaves, tree.Value)
		return leaves
	}
	for _, child := range tree.Children {
		leaves = append(leaves, leafNodes(&child)...)
	}
	return leaves
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
