package three

import (
	"adventOfCode/shared"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Mul struct {
	left  int
	right int
}

// get ready, unlicensed use of regex incoming!
func findMultiplicationPairs(input string) []Mul {
	r := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	matches := r.FindAllString(input, -1)

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

func sumMultiply(pairs []Mul) int {
	total := 0
	for _, pair := range pairs {
		total += pair.left * pair.right
	}
	return total
}

func One() {
	input := shared.ReadFile("three/input")
	pairs := findMultiplicationPairs(input)
	total := sumMultiply(pairs)
	fmt.Println(total)
}
