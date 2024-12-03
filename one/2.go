package one

import "fmt"

// if left number is the same, similarity number is also the same
// this can be used to optimise this, but on only 1000 rows, I don't care.
// with numbers up to 10^6 with 1000 rows the duplicate rate should actually be quite low,
// so maybe that kidn of optimisation would be pretty useless
func findSimilarityByOccurances(leftList []int, rightList []int) int {
	totalSimilarity := 0
	for _, left := range leftList {

		occurances := 0
		for _, right := range rightList {
			if left == right {
				occurances += 1
			}
		}

		totalSimilarity += left * occurances
	}

	return totalSimilarity
}

func Two() {
	leftList, rightList := getListsFromInput()

	similarity := findSimilarityByOccurances(leftList, rightList)
	fmt.Println(similarity)
}
