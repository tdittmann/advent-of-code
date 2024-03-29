package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
	"slices"
	"strings"
)

func main() {

	day := "04"

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

	for _, line := range puzzleInput {

		points := 0

		numberString := strings.Split(line, ":")[1]
		splittedNumberString := strings.Split(numberString, "|")

		winningNumbers := findNumbers(strings.Split(splittedNumberString[0], " "))
		numbersToCheck := findNumbers(strings.Split(splittedNumberString[1], " "))

		for _, number := range numbersToCheck {
			if slices.Contains(winningNumbers, number) {
				if points == 0 {
					points++
				} else {
					points *= 2
				}
			}
		}

		result += points
	}

	return result
}

func findNumbers(input []string) []string {

	var result []string

	for _, character := range input {
		if character != "" {
			result = append(result, strings.TrimSpace(character))
		}
	}

	return result
}

func SolvePartTwo(puzzleInput []string) int {

	result := 0

	gameMap := make(map[int]int)

	for _, line := range puzzleInput {

		// First extract the game number and fill the map
		gameNumber := extractGameNumber(line)
		_, existing := gameMap[gameNumber]
		if !existing {
			gameMap[gameNumber] = 1
		}

		// For each copies we need to increment the copies of the other cards
		numberOfCopies := gameMap[gameNumber]
		for i := 0; i < numberOfCopies; i++ {

			numberString := strings.Split(line, ":")[1]
			splittedNumberString := strings.Split(numberString, "|")
			winningNumbers := findNumbers(strings.Split(splittedNumberString[0], " "))
			numbersToCheck := findNumbers(strings.Split(splittedNumberString[1], " "))
			counter := 0
			for _, number := range numbersToCheck {
				if slices.Contains(winningNumbers, number) {
					counter++
				}
			}

			// Add the copies to the map
			for i := 1; i <= counter; i++ {

				gameNumberToIncrement := gameNumber + i
				_, existing = gameMap[gameNumberToIncrement]
				if !existing {
					gameMap[gameNumberToIncrement] = 1
				}

				gameMap[gameNumberToIncrement]++
			}
		}
	}

	// Sum up the copies
	for _, value := range gameMap {
		result += value
	}

	return result
}

func extractGameNumber(input string) int {

	return helpers.ExtractNumber(strings.Split(input, ":")[0])
}
