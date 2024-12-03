package two

import (
	"adventOfCode/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func countValidReports(reports [][]int) int {
	totalValid := 0
	for range reports {

	}
	return totalValid

}

func parseLines(lines []string) [][]int {
	lineCount := len(lines)

	reports := make([][]int, lineCount, lineCount)

	for index, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			break
		}

		levels := strings.Split(line, " ")

		levelCount := len(levels)

		reports[index] = make([]int, levelCount, levelCount)
		for levelIndex, level := range levels {
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("Failed to read line")
			}

			reports[index][levelIndex] = levelInt
		}
	}

	return reports
}

func One() {
	lines := shared.ReadFileLines("two/input")
	reports := parseLines(lines)
	validReports := countValidReports(reports)
	fmt.Println(validReports)
}
