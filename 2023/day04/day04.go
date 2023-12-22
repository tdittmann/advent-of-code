package day04

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

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

	return extractNumber(strings.Split(input, ":")[0])
}

func extractNumber(input string) int {
	numberRegex := regexp.MustCompile("[0-9]+")
	number, _ := strconv.Atoi(numberRegex.FindAllString(input, -1)[0])
	return number
}
