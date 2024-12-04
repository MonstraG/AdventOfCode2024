package three

import (
	"reflect"
	"testing"
)

func TestFindMultiplicationPairs(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got := findMultiplicationPairs(input)
	want := []Mul{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
