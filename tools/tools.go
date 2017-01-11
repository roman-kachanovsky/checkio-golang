package tools

// Where are my generics? :/

func MaxFromPairOfInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinFromPairOfInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func SumInt(a []int) int {
	var s int

	for _, v := range a {
		s += v
	}
	return s
}

func MaxFloat64(a []float64) float64 {
	if len(a) == 0 {
		return 0
	}

	var biggest float64 = a[0]

	for _, v := range a {
		if biggest < v {
			biggest = v
		}
	}
	return biggest
}

func MinFloat64(a []float64) float64 {
	if len(a) == 0 {
		return 0
	}

	var smallest float64 = a[0]

	for _, v := range a {
		if smallest > v {
			smallest = v
		}
	}
	return smallest
}
