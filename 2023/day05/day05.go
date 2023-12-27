package main

import (
	"adventofcode2023/2023/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	day := "05"

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

	var seeds []int
	var mappings []mapping
	currentMapping := mapping{}

	for _, line := range puzzleInput {

		// Is it an empty line? Then ignore it
		if line == "" {
			continue
		}

		// If it's the first line, extract all the seeds
		if strings.Contains(line, "seeds:") {
			splittedSeeds := strings.Split(line, "seeds: ")
			splittedNumbers := strings.Split(splittedSeeds[1], " ")
			for _, seedNumber := range splittedNumbers {
				seeds = append(seeds, extractNumber(seedNumber))
			}
			continue
		}

		// If it's a new mapping line, we need to close the current mapping
		if strings.Contains(line, ":") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = mapping{}
			continue
		}

		values := strings.Split(line, " ")

		// Add the mapping with starting index, size and offset
		currentMapping = append(currentMapping, submapping{
			sourceRangeStart: extractNumber(values[1]),
			size:             extractNumber(values[2]),
			offset:           extractNumber(values[0]) - extractNumber(values[1]),
		})
	}

	// If we have some mappings left, add them to the slice
	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	// Search the lowest location number
	result := 99999999999999
	for _, seed := range seeds {

		for _, mapping := range mappings {
			for _, submapping := range mapping {
				if seed >= submapping.sourceRangeStart && seed <= submapping.sourceRangeStart+submapping.size {
					seed += submapping.offset
					break
				}
			}
		}

		result = min(result, seed)
	}

	return result
}

type mapping []submapping

type submapping struct {
	sourceRangeStart, size, offset int
}

func extractNumber(input string) int {
	numberRegex := regexp.MustCompile("[0-9]+")
	number, _ := strconv.Atoi(numberRegex.FindAllString(input, -1)[0])
	return number
}

func SolvePartTwo(puzzleInput []string) int {

	result := 0

	return result
}
