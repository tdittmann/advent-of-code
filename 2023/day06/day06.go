package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
	"strings"
)

func main() {

	day := "06"

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

	// First exract all times and distances
	var times []int
	var distanceRecords []int

	for _, line := range puzzleInput {

		if strings.Contains(line, "Time:") {
			splittedString := strings.Split(line, "Time:")
			timesString := strings.TrimSpace(splittedString[1])
			splittedTimes := strings.Split(timesString, " ")
			for _, time := range splittedTimes {
				if len(time) > 0 {
					times = append(times, helpers.ExtractNumber(time))
				}
			}
		}

		if strings.Contains(line, "Distance:") {
			splittedString := strings.Split(line, "Distance:")
			timesString := strings.TrimSpace(splittedString[1])
			splittedTimes := strings.Split(timesString, " ")
			for _, time := range splittedTimes {
				if len(time) > 0 {
					distanceRecords = append(distanceRecords, helpers.ExtractNumber(time))
				}
			}
		}
	}

	result := 1

	for i, time := range times {

		numberOfDifferentWaysToWin := 0
		distanceRecord := distanceRecords[i]

		for j := 0; j < time; j++ {
			speed := j
			travelTime := time - j
			distance := travelTime * speed

			if distance > distanceRecord {
				numberOfDifferentWaysToWin++
			}
		}

		result *= numberOfDifferentWaysToWin
	}

	return result
}

func SolvePartTwo(puzzleInput []string) int {

	var time int
	var distanceRecord int

	for _, line := range puzzleInput {

		if strings.Contains(line, "Time:") {
			splittedString := strings.Split(line, "Time:")
			timesString := strings.TrimSpace(splittedString[1])
			splittedTimes := strings.ReplaceAll(timesString, " ", "")
			time = helpers.ExtractNumber(splittedTimes)
		}

		if strings.Contains(line, "Distance:") {
			splittedString := strings.Split(line, "Distance:")
			timesString := strings.TrimSpace(splittedString[1])
			splitted := strings.ReplaceAll(timesString, " ", "")
			distanceRecord = helpers.ExtractNumber(splitted)
		}
	}

	result := 0

	for j := 0; j < time; j++ {
		speed := j
		travelTime := time - j
		distance := travelTime * speed

		if distance > distanceRecord {
			result++
		}
	}

	return result
}
