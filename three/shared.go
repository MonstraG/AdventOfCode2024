package three

import (
	"regexp"
	"strconv"
	"strings"
)

func sumMultiply(pairs []Mul) int {
	total := 0
	for _, pair := range pairs {
		total += pair.left * pair.right
	}
	return total
}

type Mul struct {
	left  int
	right int
}

// ["mul(1,2)", "mul(3,4)"] -> [{left: 1, right: 2}, {left: 3, right: 4}]
func parseMatches(matches []string) []Mul {
	matchCount := len(matches)

	result := make([]Mul, matchCount, matchCount)
	for index, match := range matches {
		trimmed := strings.Trim(match, "mul()")
		parts := strings.Split(trimmed, ",")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Failed to parse left in" + trimmed)
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Failed to parse left in" + trimmed)
		}
		result[index] = Mul{left, right}
	}

	return result
}

var mulStatementRegex = regexp.MustCompile("mul\\(\\d+,\\d+\\)")

func findAllMuls(input string) []string {
	return mulStatementRegex.FindAllString(input, -1)
}
