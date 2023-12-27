package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadContentOfFile(path string) []string {

	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return splitStringByNewLine(string(dat))
}

func splitStringByNewLine(input string) []string {

	return strings.Split(input, "\n")
}

func ExtractNumber(input string) int {
	numberRegex := regexp.MustCompile("[0-9]+")
	number, _ := strconv.Atoi(numberRegex.FindAllString(input, -1)[0])
	return number
}
