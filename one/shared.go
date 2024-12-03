package one

import (
	"adventOfCode/shared"
	"log"
	"strconv"
	"strings"
)

func parseLists(file string) ([]int, []int) {
	lines := strings.Split(file, "\n")

	lineCount := len(lines)

	leftList := make([]int, lineCount, lineCount)
	rightList := make([]int, lineCount, lineCount)

	var err error

	for lineIndex, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			// assume end
			break
		}

		numbers := strings.Split(trimmed, "   ")

		if len(numbers) != 2 {
			log.Println("Weird line found " + strconv.Itoa(lineIndex))
			log.Println(trimmed)
			panic("fail")
		}

		leftList[lineIndex], err = strconv.Atoi(numbers[0])
		if err != nil {
			log.Print("Failed to left parse number in line " + strconv.Itoa(lineIndex))
			panic(err)
		}

		rightList[lineIndex], err = strconv.Atoi(numbers[1])
		if err != nil {
			log.Print("Failed to right parse number in line " + strconv.Itoa(lineIndex))
			panic(err)
		}
	}

	return leftList, rightList
}

func getListsFromInput() ([]int, []int) {
	file := shared.ReadFile()
	return parseLists(file)
}
