package elementary

import "testing"

func TestEvenTheLast(t *testing.T) {
	cases := []struct {
		in []int
		want int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 30},
		{[]int{1, 3, 5}, 30},
		{[]int{6}, 36},
		{[]int{}, 0},
	}

	for _, c := range cases {
		got := even_the_last(c.in)
		if got != c.want {
			t.Errorf("even_the_last(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
