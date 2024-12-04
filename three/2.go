package three

import (
	"adventOfCode/shared"
	"fmt"
	"strings"
)

const doMode = "do()"
const dontMode = "don't()"

func findEnabledMemoryBlocks(input string) []string {
	currentMode := doMode

	enabledInputParts := make([]string, 0)

	copy := input
	for len(copy) > 0 {

		// figure out what is the next switch that will have effect
		// (2 of the same switch do nothing)
		var modeSwitch string
		if currentMode == doMode {
			modeSwitch = dontMode
		} else {
			modeSwitch = doMode
		}

		// split by that switch
		parts := strings.SplitN(copy, modeSwitch, 2)
		left := parts[0]

		// if we are in enabled mode left part is enabled
		if currentMode == doMode {
			enabledInputParts = append(enabledInputParts, left)
		}

		// toggle the mode lol
		currentMode = modeSwitch

		// put remainder back
		hasRemainder := len(parts) > 1
		if hasRemainder {
			copy = parts[1]
		} else {
			return enabledInputParts
		}
	}

	return enabledInputParts
}

// get ready, unlicensed use of regex incoming!
func findConditionalMultiplicationPairs(input string) []Mul {
	enabledBlocks := findEnabledMemoryBlocks(input)

	operations := make([]Mul, 0)
	for _, block := range enabledBlocks {
		matches := findAllMuls(block)
		operations = append(operations, parseMatches(matches)...)
	}

	return operations
}

func Two() {
	input := shared.ReadFile("three/input")
	pairs := findConditionalMultiplicationPairs(input)
	total := sumMultiply(pairs)
	fmt.Println(total)
}
