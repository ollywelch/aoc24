package day1_test

import (
	"testing"

	"github.com/ollywelch/aoc24/day1"
)

func TestCalculateDistance(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate distance of 0",
			args: args{
				nums1: []int{2, 1, 3},
				nums2: []int{3, 1, 2},
			},
			want: 0,
		},
		{
			name: "Should calculate distance of 10",
			args: args{
				nums1: []int{2, 1, 13},
				nums2: []int{3, 1, 2},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.CalculateDistance(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSimilarity(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate given example correctly",
			args: args{
				nums1: []int{3, 4, 2, 1, 3, 3},
				nums2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
		{
			name: "Should calculate another example correctly",
			args: args{
				nums1: []int{3, 4, 2},
				nums2: []int{4, 3, 5},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.CalculateSimilarity(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("CalculateSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
