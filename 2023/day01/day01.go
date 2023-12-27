package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	day := "01"

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

		result += findNumberToAdd(line)

	}

	return result
}

func findNumberToAdd(input string) int {

	re := regexp.MustCompile("[0-9]+")
	allNumbers := re.FindAllString(input, -1)

	if len(allNumbers) == 0 {
		return 0
	}

	firstNumber := allNumbers[0]
	firstNumberLastDigit := string(firstNumber[0])
	secondNumber := allNumbers[len(allNumbers)-1]
	secondNumberLastDigit := string(secondNumber[len(secondNumber)-1])

	numberToAdd, _ := strconv.Atoi(firstNumberLastDigit + secondNumberLastDigit)

	return numberToAdd
}

func SolvePartTwo(puzzleInput []string) int {

	result := 0

	for _, line := range puzzleInput {

		firstDigit := 0
		firstSet := false
		lastDigit := 0

		for i, char := range line {

			var dig int
			found := false

			if char >= '0' && char <= '9' {
				dig = int(char - '0')
				found = true
			}

			if !found {
				switch {
				case checkWord(line, i, "one"):
					found = true
					dig = 1
				case checkWord(line, i, "two"):
					found = true
					dig = 2
				case checkWord(line, i, "three"):
					found = true
					dig = 3
				case checkWord(line, i, "four"):
					found = true
					dig = 4
				case checkWord(line, i, "five"):
					found = true
					dig = 5
				case checkWord(line, i, "six"):
					found = true
					dig = 6
				case checkWord(line, i, "seven"):
					found = true
					dig = 7
				case checkWord(line, i, "eight"):
					found = true
					dig = 8
				case checkWord(line, i, "nine"):
					found = true
					dig = 9
				}
			}

			if !found {
				continue
			}

			if !firstSet {
				firstDigit = dig
				firstSet = true
			}

			lastDigit = dig
		}

		result += (firstDigit * 10) + lastDigit
	}

	return result
}

func checkWord(line string, i int, word string) bool {
	for j := range word {
		if i+j >= len(line) {
			return false
		}

		if line[i+j] != word[j] {
			return false
		}
	}

	return true
}
