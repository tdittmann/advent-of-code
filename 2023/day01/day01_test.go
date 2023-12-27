package main

import (
	"adventofcode2023/2023/helpers"
	"testing"
)

func TestPartOne(t *testing.T) {

	var input = helpers.ReadContentOfFile("test1.txt")

	var actual = SolvePartOne(input)

	var expected = 142
	if actual != expected {
		t.Errorf("Expected [%d] but actual is [%d]", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {

	var input = helpers.ReadContentOfFile("test2.txt")

	var actual = SolvePartTwo(input)

	var expected = 281
	if actual != expected {
		t.Errorf("Expected [%d] but actual is [%d]", expected, actual)
	}
}
