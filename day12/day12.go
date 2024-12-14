package day12

import (
	"fmt"
	"slices"
)

type Region map[int]map[int]int

type Plot struct {
	Region        Region
	RightBoundary Region
	LeftBoundary  Region
	UpBoundary    Region
	DownBoundary  Region
	Char          byte
}

func NewPlot(char byte) Plot {
	return Plot{
		Region:        make(Region),
		RightBoundary: make(Region),
		LeftBoundary:  make(Region),
		UpBoundary:    make(Region),
		DownBoundary:  make(Region),
		Char:          char,
	}
}

func CalculatePrice2(grid []string) int {
	total := 0
	plots := search(grid)
	for _, plot := range plots {
		regionArea := area(plot.Region)
		rowEdges := countRowEdges(transpose(plot.RightBoundary)) + countRowEdges(transpose(plot.LeftBoundary))
		columnEdges := countRowEdges(plot.UpBoundary) + countRowEdges(plot.DownBoundary)
		calc := regionArea * (rowEdges + columnEdges)
		fmt.Println(string(plot.Char), regionArea, rowEdges, columnEdges, calc)
		total += calc
	}
	return total
}

func CalculatePrice(grid []string) int {
	// strategy
	// - go through each element in the grid
	// - for each letter, group those that are in the same region
	// - calculate length of perimeter?
	// - count area
	// - multiply and add to cumulative total

	total := 0
	plots := search(grid)

	for _, plot := range plots {
		regionArea := area(plot.Region)
		horizArea := area(plot.UpBoundary) + area(plot.DownBoundary)
		vertArea := area(plot.LeftBoundary) + area(plot.RightBoundary)
		calc := regionArea * (horizArea + vertArea)
		total += calc
	}

	return total
}

func search(grid []string) []Plot {
	visited := Region{}
	plots := []Plot{}

	for i := range grid {
		for j := range grid[i] {
			// start a walk from this point if we haven't already put the i,j in a region
			// keep walking
			if !inRegion(visited, i, j) {
				plot := NewPlot(grid[i][j])
				plot = walk(grid, plot, i, j)
				visited = merge(visited, plot.Region)
				plots = append(plots, plot)
			}
		}
	}

	return plots
}

// collects up all the coordinates in a region starting from i,j
func walk(grid []string, visited Plot, i, j int) Plot {
	char := grid[i][j]
	visitedRow, ok := visited.Region[i]
	if !ok {
		visitedRow = make(map[int]int)
	}
	visitedRow[j] = 1
	visited.Region[i] = visitedRow
	if i > 0 {
		up := grid[i-1][j]
		if up == char && !inRegion(visited.Region, i-1, j) {
			upPlot := walk(grid, visited, i-1, j)
			visited.Region = merge(upPlot.Region, visited.Region)
			visited.UpBoundary = merge(upPlot.UpBoundary, visited.UpBoundary)
		}
		if up != char {
			visited.UpBoundary = insert(visited.UpBoundary, i-1, j)
		}
	} else {
		visited.UpBoundary = insert(visited.UpBoundary, i-1, j)
	}
	if i < len(grid)-1 {
		down := grid[i+1][j]
		if down == char && !inRegion(visited.Region, i+1, j) {
			downPlot := walk(grid, visited, i+1, j)
			visited.Region = merge(downPlot.Region, visited.Region)
			visited.DownBoundary = merge(downPlot.DownBoundary, visited.DownBoundary)
		}
		if down != char {
			visited.DownBoundary = insert(visited.DownBoundary, i+1, j)
		}
	} else {
		visited.DownBoundary = insert(visited.DownBoundary, i+1, j)
	}
	if j > 0 {
		left := grid[i][j-1]
		if left == char && !inRegion(visited.Region, i, j-1) {
			leftPlot := walk(grid, visited, i, j-1)
			visited.Region = merge(leftPlot.Region, visited.Region)
			visited.LeftBoundary = merge(leftPlot.LeftBoundary, visited.LeftBoundary)
		}
		if left != char {
			visited.LeftBoundary = insert(visited.LeftBoundary, i, j-1)
		}
	} else {
		visited.LeftBoundary = insert(visited.LeftBoundary, i, j-1)
	}
	if j < len(grid[0])-1 {
		right := grid[i][j+1]
		if right == char && !inRegion(visited.Region, i, j+1) {
			rightPlot := walk(grid, visited, i, j+1)
			visited.Region = merge(rightPlot.Region, visited.Region)
			visited.RightBoundary = merge(rightPlot.RightBoundary, visited.RightBoundary)
		}
		if right != char {
			visited.RightBoundary = insert(visited.RightBoundary, i, j+1)
		}
	} else {
		visited.RightBoundary = insert(visited.RightBoundary, i, j+1)
	}
	return visited
}

func area(r Region) int {
	total := 0
	for i := range r {
		for j := range r[i] {
			total += r[i][j]
		}
	}
	return total
}

func countRowEdges(r Region) int {
	edges := 0
	for row := range r {
		breaks := 1
		cols := []int{}
		weights := []int{}
		for col := range r[row] {
			cols = append(cols, col)
			weights = append(weights, r[row][col])
		}
		slices.Sort(cols)
		for i := range cols {
			if i == 0 {
				continue
			}
			if cols[i]-1 != cols[i-1] || weights[i] != weights[i-1] {
				breaks += weights[i]
			}
		}
		edges += breaks
	}
	return edges
}

func transpose(r Region) Region {
	transposed := make(Region)
	for i := range r {
		for j := range r[i] {
			row, ok := transposed[j]
			if !ok {
				row = make(map[int]int)
			}
			row[i] = r[i][j]
			transposed[j] = row
		}
	}
	return transposed
}

func insert(r Region, i, j int) Region {
	row, ok := r[i]
	if !ok {
		row = make(map[int]int)
	}
	val, ok := row[j]
	if !ok {
		val = 0
	}
	val++
	row[j] = val
	r[i] = row
	return r
}

func inRegion(r Region, i, j int) bool {
	row, ok := r[i]
	if !ok {
		return false
	}
	_, ok = row[j]
	return ok
}

func merge(r1, r2 Region) Region {
	merged := r1
	for i := range r2 {
		row, ok := merged[i]
		if !ok {
			row = make(map[int]int)
		}
		for j := range r2[i] {
			row[j] = r2[i][j]
		}
		merged[i] = row
	}
	return merged
}
