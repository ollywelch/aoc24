package day13_test

import (
	"testing"

	"github.com/ollywelch/aoc24/day13"
)

func TestCountTokens(t *testing.T) {
	type args struct {
		machines []day13.Machine
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate correct number of tokens",
			args: args{
				machines: []day13.Machine{
					{
						ButtonA: day13.Vector{
							X: 94,
							Y: 34,
						},
						ButtonB: day13.Vector{
							X: 22,
							Y: 67,
						},
						Prize: day13.Vector{
							X: 8400,
							Y: 5400,
						},
					},
				},
			},
			want: 280,
		},
		{
			name: "Should calculate correct number of tokens",
			args: args{
				machines: []day13.Machine{
					{
						ButtonA: day13.Vector{
							X: 26,
							Y: 66,
						},
						ButtonB: day13.Vector{
							X: 67,
							Y: 21,
						},
						Prize: day13.Vector{
							X: 12748,
							Y: 12176,
						},
					},
				},
			},
			want: 0,
		},
		{
			name: "Should calculate correct number of tokens",
			args: args{
				machines: []day13.Machine{
					{
						ButtonA: day13.Vector{
							X: 17,
							Y: 86,
						},
						ButtonB: day13.Vector{
							X: 84,
							Y: 37,
						},
						Prize: day13.Vector{
							X: 7870,
							Y: 6450,
						},
					},
				},
			},
			want: 200,
		},
		{
			name: "Should calculate correct number of tokens",
			args: args{
				machines: []day13.Machine{
					{
						ButtonA: day13.Vector{
							X: 69,
							Y: 23,
						},
						ButtonB: day13.Vector{
							X: 27,
							Y: 71,
						},
						Prize: day13.Vector{
							X: 18641,
							Y: 10279,
						},
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day13.CountTokens(tt.args.machines, 100); got != tt.want {
				t.Errorf("CountTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
