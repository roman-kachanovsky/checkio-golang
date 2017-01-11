package elementary

import "testing"

func TestThreeWords(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello World hello", true},
		{"He is 123 man", false},
		{"1 2 3 4", false},
		{"bla bla bla bla", true},
		{"Hi", false},
	}

	for _, c := range cases {
		got := three_words(c.in)
		if got != c.want {
			t.Errorf("three_words(%s) == %v, want %v", c.in, got, c.want)
		}
	}
}

