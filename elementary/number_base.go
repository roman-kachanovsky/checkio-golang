package elementary

import "strconv"

func number_base(s string, base int) int {
	n, err := strconv.ParseInt(s, base, 32)

	if err != nil {
		return -1
	}

	return int(n)
}
