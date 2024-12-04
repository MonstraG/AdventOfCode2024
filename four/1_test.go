package four

import (
	"strings"
	"testing"
)

func TestXmas(t *testing.T) {
	input := `..X...
.SAMX.
.A..A.
XMAS.S
.X....`
	got := xmas(strings.Split(input, "\n"))
	want := 4
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestXmas2(t *testing.T) {
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
	got := xmas(strings.Split(input, "\n"))
	want := 18
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
