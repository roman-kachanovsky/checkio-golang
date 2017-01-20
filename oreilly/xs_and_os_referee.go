package oreilly

func xs_and_os_referee(results []string) rune {
	var x string
	combos := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, r := range results {
		x += r
	}

	for _, c := range combos {
		i, j, k := c[0], c[1], c[2]

		if x[i] == x[j] && x[j] == x[k] && (x[i] == 88 || x[i] == 79) {
			return rune(x[i])
		}
	}

	return 'D'
}
