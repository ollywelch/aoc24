package day10_test

import (
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day10"
)

func TestCalculateScore(t *testing.T) {
	example := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	exampleGrid := strings.Split(example, "\n")
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate score",
			args: args{
				grid: exampleGrid,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10.CalculateScore(tt.args.grid); got != tt.want {
				t.Errorf("CalculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRating(t *testing.T) {
	example := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	exampleGrid := strings.Split(example, "\n")
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate score",
			args: args{
				grid: exampleGrid,
			},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10.CalculateRating(tt.args.grid); got != tt.want {
				t.Errorf("CalculateRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
