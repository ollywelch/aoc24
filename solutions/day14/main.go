package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ollywelch/aoc24/day14"
)

func main() {
	input, err := os.Open("inputs/day14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	robots := []day14.Robot{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		robot := day14.ParseRobot(text)
		robots = append(robots, robot)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", day14.CalculateSafety(robots, 103, 101, 100))

	day14.DisplayRobots(robots, 103, 101, 0, 2000)
}
