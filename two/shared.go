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

// getReportDirection goes trough entire report and figures out what is the
// "majority" direction for the report:
// 3, 1, 2, 3, -> +
// 8, 7, 6, 7  -> -
func getReportDirection(report []int) int {
	direction := 0
	prev := report[0]
	for _, level := range report {
		levelDiff := diff(prev, level)
		if levelDiff > 0 {
			direction += 1
		}
		if levelDiff < 0 {
			direction -= 1
		}
	}

	// and clamp to [-1, 1]
	if direction > 0 {
		return 1
	}
	if direction < 0 {
		return -1
	}
	return 0
}

// diff is just the minus operation
// the only reason it's a separate function is to make sure that you pass
// stuff in the correct order, meaning it's always a - b, and not b - a.
func diff(previous int, current int) int {
	return previous - current
}

func isReportValid(report []int) bool {
	if len(report) < 2 {
		// this is unspecified, but I'll assume that report of 1 number or less is invalid
		return false
	}

	reportDiff := getReportDirection(report)
	if !isDifferenceValid(reportDiff, reportDiff) {
		return false
	}

	previous := report[0]
	for _, level := range report[1:] {
		levelDiff := diff(previous, level)
		if !isDifferenceValid(reportDiff, levelDiff) {
			return false
		}

		previous = level
	}

	return true
}
