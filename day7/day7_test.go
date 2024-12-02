package day7_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day7"
)

func TestIsValid(t *testing.T) {
	inputs := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	splitInputs := strings.Split(inputs, "\n")

	targets := []int{}
	seqs := [][]int{}
	for _, input := range splitInputs {
		before, after, _ := strings.Cut(input, ": ")
		target, err := strconv.Atoi(before)
		if err != nil {
			t.Fatal(err)
		}
		targets = append(targets, target)

		splitNums := strings.Split(after, " ")
		seq := []int{}
		for _, strNum := range splitNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				t.Fatal(err)
			}
			seq = append(seq, num)
		}
		seqs = append(seqs, seq)
	}

	type args struct {
		target int
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should determine validity 1",
			args: args{
				target: targets[0],
				inputs: seqs[0],
			},
			want: true,
		},
		{
			name: "should determine validity 2",
			args: args{
				target: targets[1],
				inputs: seqs[1],
			},
			want: true,
		},
		{
			name: "should determine validity 3",
			args: args{
				target: targets[2],
				inputs: seqs[2],
			},
			want: false,
		},
		{
			name: "should determine validity 4",
			args: args{
				target: targets[3],
				inputs: seqs[3],
			},
			want: false,
		},
		{
			name: "should determine validity 5",
			args: args{
				target: targets[4],
				inputs: seqs[4],
			},
			want: false,
		},
		{
			name: "should determine validity 6",
			args: args{
				target: targets[5],
				inputs: seqs[5],
			},
			want: false,
		},
		{
			name: "should determine validity 7",
			args: args{
				target: targets[6],
				inputs: seqs[6],
			},
			want: false,
		},
		{
			name: "should determine validity 8",
			args: args{
				target: targets[7],
				inputs: seqs[7],
			},
			want: false,
		},
		{
			name: "should determine validity 9",
			args: args{
				target: targets[8],
				inputs: seqs[8],
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7.IsValid(tt.args.target, tt.args.inputs); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidWithConcatenation(t *testing.T) {
	inputs := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	splitInputs := strings.Split(inputs, "\n")

	targets := []int{}
	seqs := [][]int{}
	for _, input := range splitInputs {
		before, after, _ := strings.Cut(input, ": ")
		target, err := strconv.Atoi(before)
		if err != nil {
			t.Fatal(err)
		}
		targets = append(targets, target)

		splitNums := strings.Split(after, " ")
		seq := []int{}
		for _, strNum := range splitNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				t.Fatal(err)
			}
			seq = append(seq, num)
		}
		seqs = append(seqs, seq)
	}
	type args struct {
		target int
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should determine validity 1",
			args: args{
				target: targets[0],
				inputs: seqs[0],
			},
			want: true,
		},
		{
			name: "should determine validity 2",
			args: args{
				target: targets[1],
				inputs: seqs[1],
			},
			want: true,
		},
		{
			name: "should determine validity 3",
			args: args{
				target: targets[2],
				inputs: seqs[2],
			},
			want: false,
		},
		{
			name: "should determine validity 4",
			args: args{
				target: targets[3],
				inputs: seqs[3],
			},
			want: true,
		},
		{
			name: "should determine validity 5",
			args: args{
				target: targets[4],
				inputs: seqs[4],
			},
			want: true,
		},
		{
			name: "should determine validity 6",
			args: args{
				target: targets[5],
				inputs: seqs[5],
			},
			want: false,
		},
		{
			name: "should determine validity 7",
			args: args{
				target: targets[6],
				inputs: seqs[6],
			},
			want: true,
		},
		{
			name: "should determine validity 8",
			args: args{
				target: targets[7],
				inputs: seqs[7],
			},
			want: false,
		},
		{
			name: "should determine validity 9",
			args: args{
				target: targets[8],
				inputs: seqs[8],
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7.IsValidWithConcatenation(tt.args.target, tt.args.inputs); got != tt.want {
				t.Errorf("IsValidWithConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
