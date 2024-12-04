package four

import (
	"adventOfCode/shared"
	"fmt"
)

type Pos struct {
	row int
	col int
}

func findStarts(lines []string) []Pos {
	xStarts := make([]Pos, 0)
	for row, line := range lines {
		for col, char := range line {
			if char == 'X' {
				xStarts = append(xStarts, Pos{row, col})
			}
		}
	}

	return xStarts
}

type Jump struct {
	row          int
	col          int
	targetLetter byte
}

type Direction []Jump

// if you know where X is, do these jumps to find M,A and S
var horizontal = Direction{{0, 1, 'M'}, {0, 2, 'A'}, {0, 3, 'S'}}
var vertical = Direction{{1, 0, 'M'}, {2, 0, 'A'}, {3, 0, 'S'}}
var diagonalDescending = Direction{{1, 1, 'M'}, {2, 2, 'A'}, {3, 3, 'S'}}
var diagonalAscending = Direction{{1, -1, 'M'}, {2, -2, 'A'}, {3, -3, 'S'}}
var horizontalReverse = Direction{{0, -1, 'M'}, {0, -2, 'A'}, {0, -3, 'S'}}
var verticalReverse = Direction{{-1, 0, 'M'}, {-2, 0, 'A'}, {-3, 0, 'S'}}
var diagonalDescendingReverse = Direction{{-1, -1, 'M'}, {-2, -2, 'A'}, {-3, -3, 'S'}}
var diagonalAscendingReverse = Direction{{-1, 1, 'M'}, {-2, 2, 'A'}, {-3, 3, 'S'}}

// looking at the coordinates it's kind of clear that this is basically
// permutations of 1,2,3 pairs, but keeping this as expicit directions is
// more clear I think.
var allDirections = []Direction{
	horizontal,
	vertical,
	diagonalDescending,
	diagonalAscending,
	horizontalReverse,
	verticalReverse,
	diagonalDescendingReverse,
	diagonalAscendingReverse,
}

func xmas(lines []string) int {

	rowCount := len(lines)
	colCount := len(lines[0])

	inBounds := func(targetRow int, targetCol int) bool {
		return targetRow >= 0 && targetRow < rowCount &&
			targetCol >= 0 && targetCol < colCount
	}

	xStarts := findStarts(lines)
	total := 0
	for _, start := range xStarts {
		for _, direction := range allDirections {
			for _, jump := range direction {
				want := jump.targetLetter

				targetRow := start.row + jump.row
				targetCol := start.col + jump.col

				if !inBounds(targetRow, targetCol) {
					break
				}

				got := lines[targetRow][targetCol]
				if want != got {
					break
				}

				if got == 'S' {
					total += 1
				}
			}
		}
	}

	return total
}

func One() {
	input := shared.ReadFileLines("four/input")
	count := xmas(input)
	fmt.Println(count)
}
