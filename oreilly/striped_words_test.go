package oreilly

import "testing"

func TestStripedWords(t *testing.T) {
	cases := []struct{
		in string
		want int
	}{
		{"My name is ...", 3},
		{"Hello world", 0},
		{"A quantity of striped words.", 1},
		{"Dog,cat,mouse,bird.Human.", 3},
		{"1st 2a ab3er root rate", 1},
	}

	for _, c := range cases {
		got := striped_words(c.in)

		if got != c.want {
			t.Errorf("striped_words(%q) == %d want %d", c.in, got, c.want)
		}
	}
}