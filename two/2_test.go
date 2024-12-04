package two

import (
	"strconv"
	"testing"
)

// TestIsDampenedReportValid covers test cases from the task
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

// TestIsDampenedReportValid2 covers test cases I made up myself
func TestIsDampenedReportValid2(t *testing.T) {
	reports := []struct {
		report []int
		valid  bool
	}{
		// should be true: first 3 can be removed.
		// at the time of writing the test, this fails because my code
		// assumses that first pair defines direction for the entire run,
		// and removes first found failure which would be 2, and is incorrect
		{[]int{3, 1, 2, 3}, true},
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
