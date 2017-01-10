package elementary

import (
	"testing"
	"reflect"
)

func TestAbsoluteSorting(t *testing.T) {
	cases := []struct {
		in []int
		want []int
	}{
		{[]int{-20, -5, 10, 15}, []int{-5, 10, 15, -20}},
		{[]int{1, 2, 3, 0}, []int{0, 1, 2, 3}},
		{[]int{-1, -2, -3, 0}, []int{0, -1, -2, -3}},
	}

	for _, c := range cases {
		got := absolute_sorting(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("absolute_sorting(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
