package codeship

import "testing"

func TestTheEndOfOther(t *testing.T) {
	cases := []struct {
		in []string
		want bool
	}{
		{[]string{"hello", "lo", "he"}, true},
		{[]string{"hello", "la", "hellow", "cow"}, false},
		{[]string{"walk", "duckwalk"}, true},
		{[]string{"one"}, false},
		{[]string{"helicopter", "li", "he"}, false},
	}

	for _, c := range cases {
		got := the_end_of_other(c.in)
		if got != c.want {
			t.Errorf("the_end_of_other(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
