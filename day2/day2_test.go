package day2_test

import (
	"testing"

	"github.com/ollywelch/aoc24/day2"
)

func TestIsSafe(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect safeness for decreasing",
			args: args{
				nums: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "Should detect unsafeness for big increases",
			args: args{
				nums: []int{1, 2, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "Should detect unsafeness for non-decreases",
			args: args{
				nums: []int{8, 6, 4, 4, 1},
			},
			want: false,
		},
		{
			name: "Should detect unsafeness for big decreases",
			args: args{
				nums: []int{9, 7, 6, 2, 1},
			},
			want: false,
		},
		{
			name: "Should detect safeness for increasing",
			args: args{
				nums: []int{1, 3, 6, 7, 9},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2.IsSafe(tt.args.nums); got != tt.want {
				t.Errorf("IsSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSafeWithDampener(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect safeness as before",
			args: args{
				nums: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "Should detect unsafeness as before if it's a lost cause",
			args: args{
				nums: []int{1, 2, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "Should detect unsafeness as before if it's a lost cause with decreasing",
			args: args{
				nums: []int{9, 7, 6, 2, 1},
			},
			want: false,
		},
		{
			name: "Should detect by removing levels",
			args: args{
				nums: []int{1, 3, 2, 4, 5},
			},
			want: true,
		},
		{
			name: "Should detect by removing levels",
			args: args{
				nums: []int{8, 6, 4, 4, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2.IsSafeWithDampener(tt.args.nums); got != tt.want {
				t.Errorf("IsSafeWithDampener() = %v, want %v", got, tt.want)
			}
		})
	}
}
