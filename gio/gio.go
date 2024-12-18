package gio

import (
	"bufio"
	"os"
	"strings"
)

func ReadLinesFromFile(file *os.File, ignoreWhitespace bool) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if ignoreWhitespace {
			line = strings.TrimSpace(line)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadLinesFromPath(path string, ignoreWhitespace bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadLinesFromFile(file, ignoreWhitespace)
}