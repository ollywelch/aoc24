package day3_test

import (
	"reflect"
	"testing"

	"github.com/ollywelch/aoc24/day3"
)

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Should parse the input correctly",
			args: args{
				input: `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			},
			want:  []int{2, 5, 11, 8},
			want1: []int{4, 5, 8, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := day3.Parse(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParse2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Should parse the input correctly",
			args: args{
				input: `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			},
			want:  []int{2, 8},
			want1: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := day3.Parse2(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Parse2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
