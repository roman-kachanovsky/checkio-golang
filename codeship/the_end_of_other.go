package codeship

import (
	"strings"
	"../tools"
)

func the_end_of_other(words []string) bool {
	var markers []bool

	for _, a := range words {
		for _, b := range words {
			if a != b {
				markers = append(markers, strings.HasSuffix(a, b))
			}
		}
	}

	return tools.AnyBool(markers)
}
