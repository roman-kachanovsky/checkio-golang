package mine

import "testing"

func TestCountInversion(t *testing.T) {
	cases := []struct {
		in []int
		want int
	}{
		{[]int{1, 2, 5, 3, 4, 7, 6}, 3},
		{[]int{0, 1, 2, 3}, 0},
		{[]int{99, -99}, 1},
		{[]int{5, 3, 2, 1, 0}, 10},
	}

	for _, c := range cases {
		got := count_inversion(c.in)
		if got != c.want {
			t.Errorf("count_inversion(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
