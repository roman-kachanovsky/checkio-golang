package codeship

import "strings"

func monkey_typing(text string, words []string) int {
	var counter int

	for _, w := range words {
		if strings.Contains(strings.ToLower(text), w) {
			counter++
		}
	}

	return counter
}
