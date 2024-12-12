package day11_test

import (
	"reflect"
	"testing"

	"github.com/ollywelch/aoc24/day11"
)

func TestDoBlink(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should go from 0 -> 1",
			args: args{
				num: 0,
			},
			want: []int{1},
		},
		{
			name: "Should go from 1 -> 2024",
			args: args{
				num: 1,
			},
			want: []int{2024},
		},
		{
			name: "Should go from 2024 -> 20,24",
			args: args{
				num: 2024,
			},
			want: []int{20, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day11.DoBlink(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoBlink() = %v, want %v", got, tt.want)
			}
		})
	}
}
