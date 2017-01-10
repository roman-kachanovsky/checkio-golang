package elementary

import "math"

func index_power(a []int, i uint) int {
	if len(a) < int(i) + 1 {
		return -1
	}
	return int(math.Pow(float64(a[i]), float64(i)))
}
