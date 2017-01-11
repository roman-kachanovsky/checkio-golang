package elementary

import "unicode"

func secret_message(s string) string {
	var letters []rune

	for _, l := range s {
		if unicode.IsUpper(l) {
			letters = append(letters, l)
		}
	}

	return string(letters)
}
