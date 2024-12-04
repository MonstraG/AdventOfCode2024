package three

import (
	"adventOfCode/shared"
	"fmt"
)

// get ready, unlicensed use of regex incoming!
func findMultiplicationPairs(input string) []Mul {
	matches := findAllMuls(input)
	return parseMatches(matches)
}

func One() {
	input := shared.ReadFile("three/input")
	pairs := findMultiplicationPairs(input)
	total := sumMultiply(pairs)
	fmt.Println(total)
}
