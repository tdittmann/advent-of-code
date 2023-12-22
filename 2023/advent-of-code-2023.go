package main

import (
	"adventofcode2023/2023/day01"
	"adventofcode2023/2023/day02"
	"adventofcode2023/2023/day03"
	"adventofcode2023/2023/helpers"
	"fmt"
)

func main() {

	fmt.Println("************************************************************************")
	fmt.Println("*                       Advent of Code 2023                            *")
	fmt.Println("************************************************************************")

	puzzle01 := "2023/day01/puzzle.txt"
	fmt.Printf("Day 1 --> Part 1: [%d] | Part 2: [%d]\n",
		day01.SolvePartOne(helpers.ReadContentOfFile(puzzle01)),
		day01.SolvePartTwo(helpers.ReadContentOfFile(puzzle01)))

	puzzle02 := "2023/day02/puzzle.txt"
	fmt.Printf("Day 2 --> Part 1: [%d] | Part 2: [%d]\n",
		day02.SolvePartOne(helpers.ReadContentOfFile(puzzle02)),
		day02.SolvePartTwo(helpers.ReadContentOfFile(puzzle02)))

	puzzle03 := "2023/day03/puzzle.txt"
	fmt.Printf("Day 3 --> Part 1: [%d] | Part 2: [%d]\n",
		day03.SolvePartOne(helpers.ReadContentOfFile(puzzle03)),
		day03.SolvePartTwo(helpers.ReadContentOfFile(puzzle03)))

}
