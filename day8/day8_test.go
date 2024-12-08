package day8_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day8"
)

func TestFindAntiNodes(t *testing.T) {
	simple := `..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......#...
..........
..........`
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []day8.Coordinate
	}{
		{
			name: "Should find antinodes in simple map",
			args: args{
				input: strings.Split(simple, "\n"),
			},
			want: []day8.Coordinate{
				{
					Row:    2,
					Column: 0,
				},
				{
					Row:    1,
					Column: 3,
				},
				{
					Row:    7,
					Column: 6,
				},
				{
					Row:    6,
					Column: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day8.FindAntiNodes(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAntiNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllAntiNodes(t *testing.T) {
	simple := `T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........`

	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should find all anti nodes in the simple example",
			args: args{
				input: strings.Split(simple, "\n"),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day8.FindAllAntiNodes(tt.args.input); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("FindAllAntiNodes() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
