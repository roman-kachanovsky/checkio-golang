package elementary

func even_the_last(a []int) int {
	var sum int

	if len(a) == 0 {
		return 0
	}

	for i, v := range a {
		if i % 2 == 0 {
			sum += v
		}
	}

	return sum * a[len(a)-1]
}
