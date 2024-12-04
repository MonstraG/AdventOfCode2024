package three

import (
	"reflect"
	"testing"
)

func TestFindConditionalMultiplicationPairs(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	got := findConditionalMultiplicationPairs(input)
	want := []Mul{
		{2, 4},
		{8, 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
