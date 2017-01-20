package oreilly

import (
	"strings"
	"unicode/utf8"
	"unicode"
)

func isWord(word string) bool {
	for _, c := range word {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func striped_words(text string) int {
	PUNCTUATION := ",.!?"
	VOWELS := "AEIOUY"
	CONSONANTS := "BCDFGHJKLMNPQRSTVWXZ"

	text = strings.ToUpper(text)

	for _, c := range PUNCTUATION {
		text = strings.Replace(text, string(c), " ", -1)
	}

	for _, c := range VOWELS {
		text = strings.Replace(text, string(c), "v", -1)
	}

	for _, c := range CONSONANTS {
		text = strings.Replace(text, string(c), "c", -1)
	}

	words := strings.Split(text, " ")
	var count int

	for _, w := range words {
		if utf8.RuneCountInString(w) > 1 && isWord(w) {
			if !strings.Contains(w, "cc") && !strings.Contains(w, "vv") {
				count++
			}
		}
	}

	return count
}
