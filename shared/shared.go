package shared

import (
	"os"
	"strings"
)

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func ReadFileLines(path string) []string {
	file := ReadFile(path)
	return strings.Split(file, "\n")
}
