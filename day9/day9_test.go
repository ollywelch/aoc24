package day9_test

import (
	"testing"

	"github.com/ollywelch/aoc24/day9"
)

func TestCalculateChecksum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate checksum correctly",
			args: args{
				input: "2333133121414131402",
			},
			want: 1928,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day9.CalculateChecksum(tt.args.input); got != tt.want {
				t.Errorf("CalculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateChecksumNew(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate checksum correctly",
			args: args{
				input: "2333133121414131402",
			},
			want: 2858,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day9.CalculateChecksumNew(tt.args.input); got != tt.want {
				t.Errorf("CalculateChecksumNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
