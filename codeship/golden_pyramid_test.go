package codeship

import "testing"

func TestGoldenPyramid(t *testing.T) {
	cases := []struct {
		in [][]int
		want int
	}{
		{[][]int{
			{1},
			{2, 3},
			{3, 3, 1},
			{3, 1, 5, 4},
			{3, 1, 3, 1, 3},
			{2, 2, 2, 2, 2, 2},
			{5, 6, 4, 5, 6, 4, 3},
		}, 23},
		{[][]int{
			{1},
			{2, 1},
			{1, 2, 1},
			{1, 2, 1, 1},
			{1, 2, 1, 1, 1},
			{1, 2, 1, 1, 1, 1},
			{1, 2, 1, 1, 1, 1, 9},
		}, 15},
		{[][]int{
			{9},
			{2, 2},
			{3, 3, 3},
			{4, 4, 4, 4},
		}, 18},
	}

	for _, c := range cases {
		got := golden_pyramid(c.in)
		if got != c.want {
			t.Errorf("golden_pyramid(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
