package one

import (
	"fmt"
	"slices"
)

func findCumulativeDifference(leftList []int, rightList []int) int {
	totalDiff := 0
	for index := range leftList {
		left := leftList[index]
		right := rightList[index]

		diff := left - right
		if diff < 0 {
			totalDiff -= diff
		} else {
			totalDiff += diff
		}
	}
	return totalDiff
}

func One() {
	leftList, rightList := getListsFromInput()

	slices.Sort(leftList)
	slices.Sort(rightList)

	difference := findCumulativeDifference(leftList, rightList)
	fmt.Println(difference)
}
