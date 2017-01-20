package oreilly

import "testing"

func TestXsAndOsReferee(t *testing.T) {
	cases := []struct{
		in []string
		want rune
	}{
		{[]string{"X.O", "XX.", "XOO"}, 'X'},
		{[]string{"OO.", "XOX", "XOX"}, 'O'},
		{[]string{"OOX", "XXO", "OXX"}, 'D'},
		{[]string{"O.X", "XX.", "XOO"}, 'X'},
	}

	for _, c := range cases {
		got := xs_and_os_referee(c.in)

		if got != c.want {
			t.Errorf("xs_and_os_referee(%q) == %q want %q", c.in, got, c.want)
		}
	}
}