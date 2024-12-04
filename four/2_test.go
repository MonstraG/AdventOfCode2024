package four

import (
	"strings"
	"testing"
)

func TestX_mas2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	got := x_mas(strings.Split(input, "\n"))
	want := 9
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
