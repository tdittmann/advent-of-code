package main

import (
	"adventofcode2023/2023/helpers"
	"testing"
)

func TestPartOne(t *testing.T) {

	var input = helpers.ReadContentOfFile("test.txt")

	var actual = SolvePartOne(input)

	var expected = 288
	if actual != expected {
		t.Errorf("Expected [%d] but actual is [%d]", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {

	var input = helpers.ReadContentOfFile("test.txt")

	var actual = SolvePartTwo(input)

	var expected = 71503
	if actual != expected {
		t.Errorf("Expected [%d] but actual is [%d]", expected, actual)
	}
}
