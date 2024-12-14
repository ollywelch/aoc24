package day12_test

import (
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day12"
)

func TestCalculatePrice(t *testing.T) {
	simpleExampleRaw := `AAAA
BBCD
BBCC
EEEC`
	simpleExample := strings.Split(simpleExampleRaw, "\n")
	exampleRaw := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	example := strings.Split(exampleRaw, "\n")

	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate correct price for simple example",
			args: args{
				grid: simpleExample,
			},
			want: 140,
		},
		{
			name: "Should calculate correct price for example",
			args: args{
				grid: example,
			},
			want: 1930,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day12.CalculatePrice(tt.args.grid); got != tt.want {
				t.Errorf("CalculatePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePrice2(t *testing.T) {
	simpleExampleRaw := `AAAA
BBCD
BBCC
EEEC`
	simpleExample := strings.Split(simpleExampleRaw, "\n")
	exampleRaw := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	example := strings.Split(exampleRaw, "\n")
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate correct price for simple example",
			args: args{
				grid: simpleExample,
			},
			want: 80,
		},
		{
			name: "Should calculate correct price for example",
			args: args{
				grid: example,
			},
			want: 1206,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day12.CalculatePrice2(tt.args.grid); got != tt.want {
				t.Errorf("CalculatePrice2() = %v, want %v", got, tt.want)
			}
		})
	}
}
