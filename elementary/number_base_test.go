package elementary

import "testing"

func TestNumberBase(t *testing.T) {
	cases := []struct {
		s string
		i int
		want int
	}{
		{"AF", 16, 175},
		{"101", 2, 5},
		{"101", 5, 26},
		{"Z", 36, 35},
		{"AB", 10, -1},
	}

	for _, c := range cases {
		got := number_base(c.s, c.i)
		if got != c.want {
			t.Errorf("number_base(%s, %d) == %d, want %d", c.s, c.i, got, c.want)
		}
	}
}
