package four

import (
	"adventOfCode/shared"
	"fmt"
)

func xmas(lines []string) int {
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

	return search(lines, 'X', allDirections)
}

func One() {
	input := shared.ReadFileLines("four/input")
	count := xmas(input)
	fmt.Println(count)
}
