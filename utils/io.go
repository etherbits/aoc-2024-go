package utils

import (
	"bufio"
	"os"
	"path"
)

func GetLines(relPath string) []string {
	lines := []string{}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path.Join(dir, relPath))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
