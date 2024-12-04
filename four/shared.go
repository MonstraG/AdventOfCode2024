package four

type Pos struct {
	row int
	col int
}

func findStarts(lines []string, targetChar rune) []Pos {
	xStarts := make([]Pos, 0)
	for row, line := range lines {
		for col, char := range line {
			if char == targetChar {
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

func search(lines []string, startLetter rune, searchDirections []Direction) int {
	rowCount := len(lines)
	colCount := len(lines[0])

	inBounds := func(targetRow int, targetCol int) bool {
		return targetRow >= 0 && targetRow < rowCount &&
			targetCol >= 0 && targetCol < colCount
	}

	xStarts := findStarts(lines, startLetter)
	total := 0
	for _, start := range xStarts {
		for _, direction := range searchDirections {
			for jumpIndex, jump := range direction {
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

				if jumpIndex == len(direction)-1 {
					total += 1
				}
			}
		}
	}

	return total
}
