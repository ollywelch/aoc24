package day8

import (
	"unicode"
)

type Coordinate struct {
	Row    int
	Column int
}

type uniqueAntiNodes map[int]map[int]struct{}

func FindAllAntiNodes(input []string) []Coordinate {
	antennaMap := locateAntennas(input)
	antiNodesMap := uniqueAntiNodes{}

	for _, coordinates := range antennaMap {
		for i, antennaA := range coordinates {
			for j, antennaB := range coordinates[:i] {
				if i == j {
					continue
				}
				// strategy is to add / subtract different integer multiples
				// of the line between two similar antennas.
				// once we go out of bounds in both directions, stop adding.
				mul := 0
				rowDiff := antennaB.Row - antennaA.Row
				colDiff := antennaB.Column - antennaA.Column
				for {
					antiNode1 := Coordinate{
						Row:    antennaA.Row,
						Column: antennaA.Column,
					}
					antiNode2 := Coordinate{
						Row:    antennaA.Row,
						Column: antennaA.Column,
					}
					antiNode1.Row += mul * rowDiff
					antiNode1.Column += mul * colDiff
					antiNode2.Row -= mul * rowDiff
					antiNode2.Column -= mul * colDiff

					newAntiNode := false
					for _, antiNode := range []Coordinate{antiNode1, antiNode2} {
						if isInBounds(antiNode, input) {
							newAntiNode = true
							insertCoordinate(antiNodesMap, antiNode)
						}
					}
					if !newAntiNode {
						break
					}
					mul += 1
				}
			}
		}
	}

	return toCoordinates(antiNodesMap)
}

func FindAntiNodes(input []string) []Coordinate {
	antennaMap := locateAntennas(input)
	antiNodesMap := uniqueAntiNodes{}
	for _, coordinates := range antennaMap {
		// find anti-node for each coordinate
		for i, antennaA := range coordinates {
			for j, antennaB := range coordinates {
				if i == j {
					continue
				}
				antiNode := Coordinate{}
				antiNode.Row = antennaA.Row - (antennaB.Row - antennaA.Row)
				antiNode.Column = antennaA.Column - (antennaB.Column - antennaA.Column)
				if isInBounds(antiNode, input) {
					antiNodesMap = insertCoordinate(antiNodesMap, antiNode)
				}
			}
		}
	}
	return toCoordinates(antiNodesMap)
}

func toCoordinates(u uniqueAntiNodes) []Coordinate {
	antiNodes := []Coordinate{}
	for row := range u {
		for column := range u[row] {
			antiNode := Coordinate{
				Row:    row,
				Column: column,
			}
			antiNodes = append(antiNodes, antiNode)
		}
	}
	return antiNodes
}

func locateAntennas(input []string) map[rune][]Coordinate {
	antennaMap := map[rune][]Coordinate{}
	for i := range input {
		for j := range input[i] {
			cell := rune(input[i][j])
			if unicode.IsLetter(cell) || unicode.IsDigit(cell) {
				antennaMap[cell] = append(antennaMap[cell], Coordinate{Row: i, Column: j})
			}
		}
	}
	return antennaMap
}

func isInBounds(c Coordinate, input []string) bool {
	return c.Row >= 0 && c.Row < len(input) && c.Column >= 0 && c.Column < len(input[0])
}

func insertCoordinate(u uniqueAntiNodes, c Coordinate) uniqueAntiNodes {
	antiNodeColumns, ok := u[c.Row]
	if !ok {
		antiNodeColumns = map[int]struct{}{}
	}
	antiNodeColumns[c.Column] = struct{}{}
	u[c.Row] = antiNodeColumns
	return u
}
