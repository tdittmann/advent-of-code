package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	day := "02"

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

	maxNumberOfRedCubes := 12
	maxNumberOfGreenCubes := 13
	maxNumberOfBlueCubes := 14

	result := 0

	for _, line := range puzzleInput {

		splittedLine := strings.Split(line, ":")

		gameNumber := extractNumber(splittedLine[0])

		splittedSubsets := strings.Split(splittedLine[1], ";")

		invalidGame := false

		for _, subset := range splittedSubsets {

			splittedColors := strings.Split(subset, ",")

			for _, color := range splittedColors {

				numberOfRedCubes := findNumberOfCubes(color, "red")
				numberOfBlueCubes := findNumberOfCubes(color, "blue")
				numberOfGreenCubes := findNumberOfCubes(color, "green")

				if numberOfRedCubes > maxNumberOfRedCubes || numberOfBlueCubes > maxNumberOfBlueCubes || numberOfGreenCubes > maxNumberOfGreenCubes {
					invalidGame = true
				}

			}

		}

		if !invalidGame {
			result += gameNumber
		}
	}

	return result
}

func SolvePartTwo(puzzleInput []string) int {

	result := 0

	for _, line := range puzzleInput {

		power := 0
		highestNumberForRedCubes := 0
		highestNumberForGreenCubes := 0
		highestNumberForBlueCubes := 0

		splittedLine := strings.Split(line, ":")
		splittedSubsets := strings.Split(splittedLine[1], ";")

		for _, subset := range splittedSubsets {

			splittedColors := strings.Split(subset, ",")

			for _, color := range splittedColors {

				numberOfRedCubes := findNumberOfCubes(color, "red")
				numberOfBlueCubes := findNumberOfCubes(color, "blue")
				numberOfGreenCubes := findNumberOfCubes(color, "green")

				if numberOfRedCubes > highestNumberForRedCubes {
					highestNumberForRedCubes = numberOfRedCubes
				}

				if numberOfGreenCubes > highestNumberForGreenCubes {
					highestNumberForGreenCubes = numberOfGreenCubes
				}

				if numberOfBlueCubes > highestNumberForBlueCubes {
					highestNumberForBlueCubes = numberOfBlueCubes
				}
			}

		}

		power = highestNumberForRedCubes * highestNumberForGreenCubes * highestNumberForBlueCubes

		result += power
	}

	return result
}

func findNumberOfCubes(input string, color string) int {

	if strings.Contains(input, color) {
		return extractNumber(input)
	}

	return 0
}

func extractNumber(input string) int {
	numberRegex := regexp.MustCompile("[0-9]+")
	number, _ := strconv.Atoi(numberRegex.FindAllString(input, -1)[0])
	return number
}
