package two

import (
	"adventOfCode/shared"
	"fmt"
)

func One() {
	lines := shared.ReadFileLines("two/input")
	reports := parseReports(lines)
	validReports := countValidReports(reports, isReportValid)
	fmt.Println(validReports)
}
