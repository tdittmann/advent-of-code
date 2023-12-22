package day01

import (
	"adventofcode2023/helpers"
	"testing"
)

func TestPartOne(t *testing.T) {

	var input = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

	var actual = SolvePartOne(helpers.SplitStringByNewLine(input))

	var expected = 142
	if actual != expected {
		t.Errorf("Expected [%d] but actual is [%d]", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {

	var input = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"

	var actual = SolvePartTwo(helpers.SplitStringByNewLine(input))

	var expected = 292
	if actual != expected {
		t.Errorf("Expected [%d] but actual is [%d]", expected, actual)
	}
}
