package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ollywelch/aoc24/day13"
)

func main() {
	input, err := os.Open("inputs/day13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	machines := []day13.Machine{}
	part2Machines := []day13.Machine{}
	var machine day13.Machine

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			machine = day13.Machine{}
			continue
		}
		if strings.HasPrefix(text, "Button A") {
			line := strings.TrimPrefix(text, "Button A: ")
			split := strings.Split(line, ", ")
			x, err := strconv.Atoi(strings.TrimPrefix(split[0], "X+"))
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(strings.TrimPrefix(split[1], "Y+"))
			if err != nil {
				log.Fatal(err)
			}
			machine.ButtonA.X = x
			machine.ButtonA.Y = y
		}
		if strings.HasPrefix(text, "Button B") {
			line := strings.TrimPrefix(text, "Button B: ")
			split := strings.Split(line, ", ")
			x, err := strconv.Atoi(strings.TrimPrefix(split[0], "X+"))
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(strings.TrimPrefix(split[1], "Y+"))
			if err != nil {
				log.Fatal(err)
			}
			machine.ButtonB.X = x
			machine.ButtonB.Y = y
		}
		if strings.HasPrefix(text, "Prize: ") {
			line := strings.TrimPrefix(text, "Prize: ")
			split := strings.Split(line, ", ")
			x, err := strconv.Atoi(strings.TrimPrefix(split[0], "X="))
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(strings.TrimPrefix(split[1], "Y="))
			if err != nil {
				log.Fatal(err)
			}
			machine.Prize.X = x
			machine.Prize.Y = y
			part2Machine := machine
			part2Machine.Prize.X = x + 10000000000000
			part2Machine.Prize.Y = y + 10000000000000
			machines = append(machines, machine)
			part2Machines = append(part2Machines, part2Machine)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", day13.CountTokens(machines, 100))
	fmt.Println("Part 2:", day13.CountTokens(part2Machines, 0))
}
