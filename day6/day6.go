package day6

const (
	DirectionUp = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

type GuardPosition struct {
	Direction Direction
	Row       int
	Column    int
}

type Direction int

// Returns the grid after tracking the guard, and a bool
// to represent whether a loop was detected (true if looped)
func TrackGuard(grid []string) ([]string, bool) {
	// Find initial guard position
	initPos := initialPosition(grid)
	positions := []GuardPosition{}
	hasLoop := false
	next := initPos

	for {
		start := next
		positions = append(positions, start)

		next = NextPosition(start, grid)

		if next.Row == start.Row && next.Column == start.Column && next.Direction == start.Direction {
			grid[start.Row] = ReplaceAtIndex(grid[start.Row], 'X', start.Column)
			break
		}

		if next.Row != start.Row || next.Column != start.Column {
			grid[start.Row] = ReplaceAtIndex(grid[start.Row], 'X', start.Column)
		}

		// check if we have a loop
		for _, pos := range positions {
			if Equal(pos, next) {
				hasLoop = true
			}
		}
		if hasLoop {
			break
		}
	}

	return grid, hasLoop
}

func initialPosition(grid []string) GuardPosition {
	position := GuardPosition{}
	for i := range grid {
		position.Row = i
		for j := range grid[i] {
			position.Column = j
			switch rune(grid[i][j]) {
			case '^':
				position.Direction = DirectionUp
				return position
			case '>':
				position.Direction = DirectionRight
				return position
			case 'v':
				position.Direction = DirectionDown
				return position
			case '<':
				position.Direction = DirectionLeft
				return position
			}
		}
	}
	panic("No initial guard position")
}

// Does one of:
// - Progress the guard to the next coordinate if there's no obstacle
// - Turn the guard's direction 90 degrees in case of obstacles
// - Nothing if they have reached the boundary of the grid
func NextPosition(p GuardPosition, grid []string) GuardPosition {
	nextRow := p.Row
	nextCol := p.Column

	switch p.Direction {
	case DirectionUp:
		nextRow = p.Row - 1
	case DirectionRight:
		nextCol = p.Column + 1
	case DirectionDown:
		nextRow = p.Row + 1
	case DirectionLeft:
		nextCol = p.Column - 1
	}

	if nextRow < 0 || nextCol < 0 || nextRow >= len(grid) || nextCol >= len(grid[nextRow]) {
		// we've reached the boundary, so do nothing
		return p
	}

	// turn the guard in case of obstacles
	if rune(grid[nextRow][nextCol]) == '#' {
		p.Direction = turn(p.Direction)
		return p
	}

	p.Row = nextRow
	p.Column = nextCol
	return p
}

func Equal(pos1, pos2 GuardPosition) bool {
	return pos1.Row == pos2.Row && pos1.Column == pos2.Column && pos1.Direction == pos2.Direction
}

func turn(d Direction) Direction {
	switch d {
	case DirectionDown:
		return DirectionLeft
	case DirectionLeft:
		return DirectionUp
	case DirectionUp:
		return DirectionRight
	case DirectionRight:
		return DirectionDown
	}
	return d
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
