package day15_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ollywelch/aoc24/day15"
)

func TestNext(t *testing.T) {
	type args struct {
		grid     []string
		position day15.Coordinate
		move     day15.Move
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 day15.Coordinate
	}{
		{
			name: "Should work out the next grid position",
			args: args{
				grid: strings.Split(`########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`, "\n"),
				position: day15.Coordinate{
					Row:    1,
					Column: 3,
				},
				move: day15.MoveRight,
			},
			want: strings.Split(`########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`, "\n"),
			want1: day15.Coordinate{
				Row:    1,
				Column: 4,
			},
		},
		{
			name: "Should work out the next grid position",
			args: args{
				grid: strings.Split(`########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`, "\n"),
				position: day15.Coordinate{
					Row:    1,
					Column: 4,
				},
				move: day15.MoveRight,
			},
			want: strings.Split(`########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`, "\n"),
			want1: day15.Coordinate{
				Row:    1,
				Column: 4,
			},
		},
		{
			name: "Should work out the next grid position",
			args: args{
				grid: strings.Split(`########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`, "\n"),
				position: day15.Coordinate{
					Row:    1,
					Column: 4,
				},
				move: day15.MoveDown,
			},
			want: strings.Split(`########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########`, "\n"),
			want1: day15.Coordinate{
				Row:    2,
				Column: 4,
			},
		},
		{
			name: "Should work out the next grid position",
			args: args{
				grid: strings.Split(`########
#....OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#...@..#
########`, "\n"),
				position: day15.Coordinate{
					Row:    6,
					Column: 4,
				},
				move: day15.MoveDown,
			},
			want: strings.Split(`########
#....OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#...@..#
########`, "\n"),
			want1: day15.Coordinate{
				Row:    6,
				Column: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := day15.Next(tt.args.grid, tt.args.position, tt.args.move)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDoMoves(t *testing.T) {
	grid, moves := day15.ParseInput(`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`)
	type args struct {
		grid  []string
		moves []day15.Move
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should calculate grid correctly after all moves",
			args: args{
				grid:  grid,
				moves: moves,
			},
			want: strings.Split(`##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`, "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day15.DoMoves(tt.args.grid, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoMoves() = \n%v, want \n%v", strings.Join(got, "\n"), strings.Join(tt.want, "\n"))
			}
		})
	}
}

func TestCalculateGPS(t *testing.T) {
	grid, moves := day15.ParseInput(`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`)
	grid = day15.DoMoves(grid, moves)
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate gps correctly",
			args: args{
				grid: grid,
			},
			want: 10092,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day15.CalculateGPS(tt.args.grid); got != tt.want {
				t.Errorf("CalculateGPS() = %v, want %v", got, tt.want)
			}
		})
	}
}
