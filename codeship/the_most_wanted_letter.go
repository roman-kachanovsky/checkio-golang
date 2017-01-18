package codeship

import "strings"

func the_most_wanted_letter(text string) rune {
	var most_wanted_letter struct{
		letter rune
		count int
	}
	var alphabet = "abcdefghijklmnopqrstunwxyz"

	for _, l := range alphabet {
		c := strings.Count(strings.ToLower(text), string(l))

		if c > most_wanted_letter.count {
			most_wanted_letter.count = c
			most_wanted_letter.letter = l
		}
	}

	return most_wanted_letter.letter
}
