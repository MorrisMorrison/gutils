package gio

import (
	"bufio"
	"os"

	"github.com/MorrisMorrison/gutils/gerror"
)

func ReadFile(path string, ignoreWhitespace bool) []string {
	file, error := os.Open(path)
	gerror.HandleError(error)
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		inputElement := scanner.Text()
		input = append(input, inputElement)
	}

	return input
}
