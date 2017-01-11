package mine

func count_inversion(sequence []int) int {
	var counter int

	for i, a := range sequence {
		for _, b := range sequence[i:] {
			if a > b {
				counter++
			}
		}
	}

	return counter
}
