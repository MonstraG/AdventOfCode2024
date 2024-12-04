package two

import (
	"adventOfCode/shared"
	"fmt"
)

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isDampenedReportValid(report []int) bool {
	ok := isReportValid(report)
	if ok {
		return true
	}

	for index := range report {
		ok = isReportValid(RemoveIndex(report, index))
		if ok {
			return true
		}
	}

	return false
}

func Two() {
	lines := shared.ReadFileLines("two/input")
	reports := parseReports(lines)
	validReports := countValidReports(reports, isDampenedReportValid)
	fmt.Println(validReports)
}
