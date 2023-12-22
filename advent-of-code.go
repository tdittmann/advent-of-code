package main

import (
	"adventofcode2023/helpers"
	"adventofcode2023/solutions/day01"
	"fmt"
)

func main() {

	fmt.Println("************************************************************************")
	fmt.Println("*                       Advent of Code 2023                            *")
	fmt.Println("************************************************************************")

	fmt.Printf("Day 1 --> Part 1: [%d] | Part 2: [%d]\n",
		day01.SolvePartOne(helpers.ReadContentOfFile("puzzles/day01.txt")),
		day01.SolvePartTwo(helpers.ReadContentOfFile("puzzles/day01.txt")),
	)

}
