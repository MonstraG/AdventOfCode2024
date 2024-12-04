package two

import (
	"strconv"
	"testing"
)

func TestIsDampenedReportValid(t *testing.T) {
	reports := []struct {
		report []int
		valid  bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}
	for index, tt := range reports {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			got := isDampenedReportValid(tt.report)
			if got != tt.valid {
				t.Errorf("got %v, want %v", got, tt.valid)
			}
		})
	}
}
