package codeship

import (
	"testing"
	"strings"
)

func TestTheMostWantedLetter(t *testing.T) {
	cases := []struct {
		in string
		want rune
	}{
		{"Hello World!", 'l'},
		{"How do you do?", 'o'},
		{"One", 'e'},
		{"Oops!", 'o'},
		{"AAaooo!!!!", 'a'},
		{"abe", 'a'},
		{"fsbd", 'b'},
		{strings.Repeat("a", 9000) + strings.Repeat("b", 1000), 'a'},
	}

	for _, c := range cases {
		got := the_most_wanted_letter(c.in)
		if got != c.want {
			t.Errorf("the_most_wanted_letter(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
