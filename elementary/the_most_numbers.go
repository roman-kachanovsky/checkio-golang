package elementary

func Min(a []float64) float64 {
	var smallest float64 = a[0]

	for _, v := range a {
		if smallest > v {
			smallest = v
		}
	}

	return smallest
}

func Max(a []float64) float64 {
	var biggest float64 = a[0]

	for _, v := range a {
		if biggest < v {
			biggest = v
		}
	}

	return biggest
}

func the_most_numbers(a []float64) float64 {
	if len(a) == 0 {
		return 0
	}
	return Max(a) - Min(a)
}
