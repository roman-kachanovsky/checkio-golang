package mine

import "testing"

func TestPearlsInTheBox(t *testing.T) {
	cases := []struct {
		s string
		n int
		want float64
	}{
		{"bbw", 3, 0.48},
		{"wwb", 3, 0.52},
		{"www", 3, 0.56},
		{"bbbb", 1, 0},
		{"wwbb", 4, 0.5},
		{"bwbwbwb", 5, 0.48},
	}

	for _, c := range cases {
		got := pearls_in_the_box(c.s, c.n)
		if got != c.want {
			t.Errorf("pearls_in_the_box(%q, %d) == %.5f, want %.2f", c.s, c.n, got, c.want)
		}
	}
}
