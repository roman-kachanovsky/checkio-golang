package elementary

import (
	"strings"
	"strconv"
)

func three_words(s string) bool {
	var words = strings.Fields(s)
	var counter uint8

	for _, w := range words {
		_, err := strconv.ParseInt(w, 10, 32)

		if err == nil {
			counter = 0
		}

		counter += 1

		if counter == 3 {
			return true
		}
	}
	return false
}
