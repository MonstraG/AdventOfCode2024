package two

import (
	"adventOfCode/shared"
	"fmt"
)

func isReportValid(report []int) bool {
	ok, _ := findReportFailure(report)
	return ok
}

func One() {
	lines := shared.ReadFileLines("two/input")
	reports := parseReports(lines)
	validReports := countValidReports(reports, isReportValid)
	fmt.Println(validReports)
}
