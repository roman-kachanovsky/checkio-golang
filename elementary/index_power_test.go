package elementary

import "testing"

func TestIndexPower(t *testing.T) {
	cases := []struct {
		a []int
		i uint
		want int
	}{
		{[]int{1, 2, 3, 4}, 2, 9},
		{[]int{1, 3, 10, 100}, 3, 1000000},
		{[]int{0, 1}, 0, 1},
		{[]int{1, 2}, 3, -1},
	}

	for _, c := range cases {
		got := index_power(c.a, c.i)
		if got != c.want {
			t.Errorf("index_power(%v, %d) == %d, want %d", c.a, c.i, got, c.want)
		}
	}
}
