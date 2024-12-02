package day5_test

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day5"
)

const rules = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`

var (
	before = []int{}
	after  = []int{}
)

func init() {
	for _, line := range strings.Split(rules, "\n") {
		splitLine := strings.Split(line, "|")
		if len(splitLine) != 2 {
			log.Fatal(fmt.Errorf("length of line %s split by | is %d", line, len(splitLine)))
		}
		b, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		a, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		before = append(before, b)
		after = append(after, a)
	}
}

func TestIsValidSequence(t *testing.T) {

	before := []int{}
	after := []int{}

	type args struct {
		seq    []int
		before []int
		after  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should spot valid sequences",
			args: args{
				seq:    []int{75, 47, 61, 53, 29},
				before: before,
				after:  after,
			},
			want: true,
		},
		{
			name: "Should spot valid sequences 2",
			args: args{
				seq:    []int{97, 61, 53, 29, 13},
				before: before,
				after:  after,
			},
			want: true,
		},
		{
			name: "Should spot valid sequences 3",
			args: args{
				seq:    []int{75, 29, 13},
				before: before,
				after:  after,
			},
			want: true,
		},
		{
			name: "Should spot invalid sequences 1",
			args: args{
				seq:    []int{75, 97, 47, 61, 53},
				before: before,
				after:  after,
			},
			want: false,
		},
		{
			name: "Should spot invalid sequences 2",
			args: args{
				seq:    []int{61, 13, 29},
				before: before,
				after:  after,
			},
			want: false,
		},
		{
			name: "Should spot invalid sequences 3",
			args: args{
				seq:    []int{97, 13, 75, 29, 47},
				before: before,
				after:  after,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day5.IsValidSequence(tt.args.seq, tt.args.before, tt.args.after); got != tt.want {
				t.Errorf("IsValidSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleNumber(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should get middle number of odd numbered lists",
			args: args{
				seq: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "Should get middle number of even numbered lists",
			args: args{
				seq: []int{3, 4, 5, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day5.MiddleNumber(tt.args.seq); got != tt.want {
				t.Errorf("MiddleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixInvalidSequence(t *testing.T) {
	type args struct {
		seq    []int
		before []int
		after  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should sort invalid sequences 1",
			args: args{
				seq:    []int{75, 97, 47, 61, 53},
				before: before,
				after:  after,
			},
			want: []int{97, 75, 47, 61, 53},
		},
		{
			name: "Should spot invalid sequences 2",
			args: args{
				seq:    []int{61, 13, 29},
				before: before,
				after:  after,
			},
			want: []int{61, 29, 13},
		},
		{
			name: "Should spot invalid sequences 3",
			args: args{
				seq:    []int{97, 13, 75, 29, 47},
				before: before,
				after:  after,
			},
			want: []int{97, 75, 47, 29, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day5.FixInvalidSequence(tt.args.seq, tt.args.before, tt.args.after); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FixInvalidSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
