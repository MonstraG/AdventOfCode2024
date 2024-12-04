package two

import (
	"adventOfCode/shared"
	"fmt"
)

// https://stackoverflow.com/a/37335777/11593686
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isDampenedReportValid(report []int) bool {
	ok, failureIndex := findReportFailure(report)
	if ok {
		return true
	}

	// remove that and try again
	ok, _ = findReportFailure(remove(report, failureIndex))
	return ok
}

func Two() {
	lines := shared.ReadFileLines("two/input")
	reports := parseReports(lines)
	validReports := countValidReports(reports, isDampenedReportValid)
	fmt.Println(validReports)
}
