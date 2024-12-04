package four

import (
	"adventOfCode/shared"
	"fmt"
)

func x_mas(lines []string) int {

	//M.S
	//.A.
	//M.S
	var diagonal1 = Direction{{-1, -1, 'M'}, {1, -1, 'S'}, {-1, 1, 'M'}, {1, 1, 'S'}}

	//S.S
	//.A.
	//M.M
	var diagonal2 = Direction{{-1, -1, 'S'}, {1, -1, 'S'}, {-1, 1, 'M'}, {1, 1, 'M'}}

	//S.M
	//.A.
	//S.M
	var diagonal3 = Direction{{-1, -1, 'S'}, {1, -1, 'M'}, {-1, 1, 'S'}, {1, 1, 'M'}}

	//M.M
	//.A.
	//S.S
	var diagonal4 = Direction{{-1, -1, 'M'}, {1, -1, 'M'}, {-1, 1, 'S'}, {1, 1, 'S'}}

	var allDirections = []Direction{
		diagonal1,
		diagonal2,
		diagonal3,
		diagonal4,
	}

	return search(lines, 'A', allDirections)
}

func Two() {
	input := shared.ReadFileLines("four/input")
	count := x_mas(input)
	fmt.Println(count)
}
