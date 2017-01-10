package elementary

import "testing"

func TestDigitsMultiplication(t *testing.T) {
	cases := []struct {
		in int
		want int
	}{
		{123405, 120},
		{999, 729},
		{1000, 1},
		{1111, 1},
	}

	for _, c := range cases {
		got := digits_multiplication(c.in)
		if got != c.want {
			t.Errorf("digits_multiplication(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
