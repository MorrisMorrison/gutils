package gio

import (
	"bufio"
	"os"
	"strconv"

	"github.com/gutils/gerror"
)

func ReadFile(path string, ignoreWhitespace bool) []string {
	file, error := os.Open(path)
	gerror.HandleError(error)
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		inputElement, error := scanner.Text())
		input = append(input, inputElement)
		gerror.HandleError(error)
	}
	
	return input
}
