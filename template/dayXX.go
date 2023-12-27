package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
)

func main() {

	day := "XX"

	fmt.Println("*************************************************************************")
	fmt.Println("*                       Advent of Code 2023                             *")
	fmt.Printf("*                              Day %s                                   *\n", day)
	fmt.Println("*************************************************************************")

	puzzleInput := fmt.Sprintf("2023/day%s/puzzle.txt", day)
	fmt.Printf("Day %s --> Part 1: [%d] | Part 2: [%d]\n",
		day,
		SolvePartOne(helpers.ReadContentOfFile(puzzleInput)),
		SolvePartTwo(helpers.ReadContentOfFile(puzzleInput)))

}

func SolvePartOne(puzzleInput []string) int {

	result := 0

	return result
}

func SolvePartTwo(puzzleInput []string) int {

	result := 0

	return result
}
