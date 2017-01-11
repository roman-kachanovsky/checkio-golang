package elementary

import "strings"

func right_to_left(a []string) string {
	var parts []string

	for _, p := range a {
		parts = append(parts, strings.Replace(p, "right", "left", -1))
	}

	return strings.Join(parts, ",")
}
