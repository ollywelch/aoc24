package day14_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day14"
)

func TestCalculateSafety(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	robots := []day14.Robot{}
	for _, line := range strings.Split(input, "\n") {
		robots = append(robots, day14.ParseRobot(line))
	}
	type args struct {
		robots  []day14.Robot
		rows    int
		columns int
		seconds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate correct safety score",
			args: args{
				robots:  robots,
				rows:    7,
				columns: 11,
				seconds: 100,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day14.CalculateSafety(tt.args.robots, tt.args.rows, tt.args.columns, tt.args.seconds); got != tt.want {
				t.Errorf("CalculateSafety() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMove(t *testing.T) {
	robot := day14.ParseRobot("p=2,4 v=2,-3")
	type args struct {
		r       day14.Robot
		rows    int
		columns int
		seconds int
	}
	tests := []struct {
		name string
		args args
		want day14.Robot
	}{
		{
			name: "Should move robots correctly",
			args: args{
				r:       robot,
				rows:    7,
				columns: 11,
				seconds: 1,
			},
			want: day14.Robot{
				Position: day14.Vector{
					X: 4,
					Y: 1,
				},
				Velocity: day14.Vector{
					X: 2,
					Y: -3,
				},
			},
		},
		{
			name: "Should move robots correctly after wrapping",
			args: args{
				r:       robot,
				rows:    7,
				columns: 11,
				seconds: 2,
			},
			want: day14.Robot{
				Position: day14.Vector{
					X: 6,
					Y: 5,
				},
				Velocity: day14.Vector{
					X: 2,
					Y: -3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day14.Move(tt.args.r, tt.args.rows, tt.args.columns, tt.args.seconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Move() = %v, want %v", got, tt.want)
			}
		})
	}
}
