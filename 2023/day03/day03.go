package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
)

func main() {

	day := "03"

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

	// Initially add all numbers to the result
	result += sumAllNumbers(puzzleInput)

	// Replace all symbols and valid numbers
	for i, line := range puzzleInput {
		for j, char := range line {
			if (char < '0' || char > '9') && char != '.' {
				for a := i - 1; a <= i+1; a++ {
					if a < 0 || a >= len(puzzleInput) {
						break
					}

					// remove all in column above in and below
					puzzleInput[a] = puzzleInput[a][:j] + "." + puzzleInput[a][j+1:]

					// remove everthing on the left side that is part of the current number
					b := j - 1
					for {
						if b < 0 || b >= len(puzzleInput[0]) {
							break
						}

						if puzzleInput[a][b] >= '0' && puzzleInput[a][b] <= '9' {
							puzzleInput[a] = puzzleInput[a][:b] + "." + puzzleInput[a][b+1:]
						} else {
							break
						}

						b--
					}

					// remove everthing on the right side that is part of the current number
					b = j + 1
					for {
						if b < 0 || b >= len(puzzleInput[0]) {
							break
						}

						if puzzleInput[a][b] >= '0' && puzzleInput[a][b] <= '9' {
							puzzleInput[a] = puzzleInput[a][:b] + "." + puzzleInput[a][b+1:]
						} else {
							break
						}

						b++
					}
				}

			}
		}
	}

	// The rest should be removed from the result
	result -= sumAllNumbers(puzzleInput)

	return result
}

func sumAllNumbers(input []string) int {

	result := 0
	buffer := 0

	for _, line := range input {
		for _, character := range line {
			if character >= '0' && character <= '9' {
				buffer *= 10
				buffer += int(character - '0')
			} else {
				result += buffer
				buffer = 0
			}
		}

		result += buffer
		buffer = 0
	}

	return result
}

func SolvePartTwo(puzzleInput []string) int {

	result := 0

	return result
}
