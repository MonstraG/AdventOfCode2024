package two

import (
	"log"
	"strconv"
	"strings"
)

func parseReports(lines []string) [][]int {
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

func countValidReports(reports [][]int, reportValidator func([]int) bool) int {
	totalValid := 0
	for _, report := range reports {
		if reportValidator(report) {
			totalValid += 1
		}
	}
	return totalValid
}

func sameSign(a, b int) bool {
	return a^b >= 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isDifferenceValid(expectedDirection int, this int) bool {
	if this == 0 {
		return false
	}
	if !sameSign(expectedDirection, this) {
		return false
	}
	if abs(this) > 3 {
		return false
	}
	return true
}

// findReportFailure
func findReportFailure(report []int) (valid bool, failureIndex int) {
	if len(report) < 2 {
		// this is unspecified, but I'll assume that report of 1 number or less is invalid
		return true, 0
	}

	previous := report[0]

	// take first difference as the canonical for report
	reportDiff := report[0] - report[1]
	if !isDifferenceValid(reportDiff, reportDiff) {
		return false, 1
	}

	for index, level := range report[1:] {
		levelDiff := previous - level
		if !isDifferenceValid(reportDiff, levelDiff) {
			return false, index + 1
		}

		previous = level
	}

	return true, 0
}
