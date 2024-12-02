package day6_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day6"
)

func TestTrackGuard(t *testing.T) {
	rawGrid := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	expectedRaw := `....#.....
....XXXXX#
....X...X.
..#.X...X.
..XXXXX#X.
..X.X.X.X.
.#XXXXXXX.
.XXXXXXX#.
#XXXXXXX..
......#X..`
	grid := strings.Split(rawGrid, "\n")
	expectedGrid := strings.Split(expectedRaw, "\n")

	type args struct {
		grid []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			name: "should record the unique coordinates of the guard in order",
			args: args{
				grid: grid,
			},
			want:  expectedGrid,
			want1: false,
		},
		{
			name: "should set the position of the first location to an X",
			args: args{
				grid: []string{
					"...",
					".^.",
					"...",
				},
			},
			want: []string{
				".X.",
				".X.",
				"...",
			},
			want1: false,
		},
		{
			name: "should successfully detect loops",
			args: args{
				grid: []string{
					".#...",
					".^..#",
					"#....",
					"...#.",
				},
			},
			want: []string{
				".#...",
				".XXX#",
				"#XXX.",
				"...#.",
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := day6.TrackGuard(tt.args.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackGuard() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TrackGuard() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
