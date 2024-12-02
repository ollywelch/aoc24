package day4_test

import (
	"reflect"
	"testing"

	"github.com/ollywelch/aoc24/day4"
)

func TestCountWords(t *testing.T) {
	type args struct {
		grid []string
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should solve xmas wordsearch",
			args: args{
				grid: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
				word: "XMAS",
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day4.CountWords(tt.args.grid, tt.args.word); got != tt.want {
				t.Errorf("CountWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountIn(t *testing.T) {
	type args struct {
		s    string
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should count forwards and backwards in a string",
			args: args{
				s:    "XMASSSAMXAAXMAS",
				word: "XMAS",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day4.CountIn(tt.args.s, tt.args.word); got != tt.want {
				t.Errorf("CountIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColumns(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should generate correct columns from a grid",
			args: args{
				grid: []string{
					"FOO",
					"BAR",
					"BAZ",
				},
			},
			want: []string{
				"FBB",
				"OAA",
				"ORZ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day4.Columns(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Columns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRDiagonals(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should generate correct LR diagonals from a grid",
			args: args{
				grid: []string{
					"FOO",
					"BAR",
					"BAZ",
					"BIZ",
				},
			},
			want: []string{
				"FAZ",
				"OR",
				"O",
				"BAZ",
				"BI",
				"B",
			},
		},
		{
			name: "Should do LR diagonals for bigger grids",
			args: args{
				grid: []string{
					"ABCDEFG",
					"HIJKLMN",
					"OPQRSTU",
					"VWXYZAB",
				},
			},
			want: []string{
				"AIQY",
				"BJRZ",
				"CKSA",
				"DLTB",
				"EMU",
				"FN",
				"G",
				"HPX",
				"OW",
				"V",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day4.LRDiagonals(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountXMAS(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should count X-MAS's",
			args: args{
				grid: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day4.CountXMAS(tt.args.grid); got != tt.want {
				t.Errorf("CountXMAS() = %v, want %v", got, tt.want)
			}
		})
	}
}
