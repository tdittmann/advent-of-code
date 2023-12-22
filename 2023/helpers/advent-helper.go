package helpers

import (
	"fmt"
	"os"
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
