package elementary

import "../tools"

func the_most_numbers(a []float64) float64 {
	if len(a) == 0 {
		return 0
	}
	return tools.MaxFloat64(a) - tools.MinFloat64(a)
}
