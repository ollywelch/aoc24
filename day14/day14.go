package day14

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Vector struct {
	X int
	Y int
}

type Robot struct {
	Position Vector
	Velocity Vector
}

func ParseRobot(input string) Robot {
	r := Robot{}

	splitText := strings.Split(input, " ")
	left := splitText[0]
	right := splitText[1]
	splitLeft := strings.Split(strings.TrimPrefix(left, "p="), ",")
	splitRight := strings.Split(strings.TrimPrefix(right, "v="), ",")

	r.Position.X = mustParseInt(splitLeft[0])
	r.Position.Y = mustParseInt(splitLeft[1])
	r.Velocity.X = mustParseInt(splitRight[0])
	r.Velocity.Y = mustParseInt(splitRight[1])
	return r
}

func DisplayRobots(robots []Robot, rows, columns, start, interval int) {
	seconds := start
	for {
		if seconds > rows*columns {
			break
		}
		robotMap := map[int]map[int]int{}
		numMiddle := 0

		for _, r := range robots {
			r = Move(r, rows, columns, seconds)
			row, ok := robotMap[r.Position.X]
			if !ok {
				row = make(map[int]int)
			}
			val, ok := row[r.Position.Y]
			if !ok {
				val = 0
			}
			val++
			row[r.Position.Y] = val
			robotMap[r.Position.X] = row
			if r.Position.Y > rows/3 && r.Position.Y < 2*rows/3 {
				numMiddle++
			}
		}

		grid := make([]string, rows)
		for i := range rows {
			gridRow := ""
			rowRobots, ok := robotMap[i]
			if ok {
				for j := range columns {
					numRobots, ok := rowRobots[j]
					if !ok {
						gridRow += "."
					} else {
						gridRow += fmt.Sprint(numRobots)
					}
				}
			} else {
				for range columns {
					gridRow += "."
				}
			}
			grid[i] = gridRow
		}
		pctMiddle := 100 * numMiddle / len(robots)
		if pctMiddle > 60 {
			fmt.Printf("Seconds: %d\n", seconds)
			fmt.Println(strings.Join(grid, "\n"))
			time.Sleep(time.Millisecond * 100)
		}
		seconds++
	}
}

func CalculateSafety(robots []Robot, rows, columns, seconds int) int {
	topLeftCount := 0
	topRightCount := 0
	bottomLeftCount := 0
	bottomRightCount := 0

	middleRow := int((rows - 1) / 2)
	middleCol := int((columns - 1) / 2)

	for _, r := range robots {
		r = Move(r, rows, columns, seconds)
		if r.Position.X < middleCol && r.Position.Y < middleRow {
			topLeftCount++
		}
		if r.Position.X > middleCol && r.Position.Y < middleRow {
			bottomLeftCount++
		}
		if r.Position.X < middleCol && r.Position.Y > middleRow {
			topRightCount++
		}
		if r.Position.X > middleCol && r.Position.Y > middleRow {
			bottomRightCount++
		}
	}
	return topLeftCount * topRightCount * bottomLeftCount * bottomRightCount
}

func Move(r Robot, rows, columns, seconds int) Robot {
	r.Position.X = (r.Position.X + (seconds * r.Velocity.X)) % columns
	r.Position.Y = (r.Position.Y + (seconds * r.Velocity.Y)) % rows
	if r.Position.X < 0 {
		r.Position.X += columns
	}
	if r.Position.Y < 0 {
		r.Position.Y += rows
	}
	return r
}

func mustParseInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
