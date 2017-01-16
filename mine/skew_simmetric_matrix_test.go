package mine

import "testing"

func TestSkewSimmetricMatrix(t *testing.T) {
	cases := []struct {
		in [][]int
		want bool
	}{
		{[][]int{{0, 1, 2}, {-1, 0, 1}, {-2, -1, 0}}, true},
		{[][]int{{0, 1, 2}, {-1, 1, 1}, {-2, -1, 0}}, false},
		{[][]int{{0, 1, 2}, {-1, 0, 1}, {-3, -1, 0}}, false},
	}

	for _, c := range cases {
		got := skew_simmetric_matrix(c.in)
		if got != c.want {
			t.Errorf("skew_simmetric_matrix(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
