package elementary

import (
	"testing"
	"math"
)

func almost_equal(checked, correct, significant_digits float64) bool {
	precision := math.Pow(0.1, significant_digits)
	if correct - precision < checked && checked < correct + precision {
		return true
	}
	return false
}

func TestTheMostNumbers(t *testing.T) {
	cases := []struct {
		in []float64
		want float64
		precision float64
	}{
		{[]float64{1, 2, 3}, 2, 3},
		{[]float64{5, -5}, 10, 3},
		{[]float64{10.2, -2.2, 0, 1.1, 0.5}, 12.4, 3},
		{[]float64{}, 0, 3},
	}

	for _, c := range cases {
		got := the_most_numbers(c.in)
		if !almost_equal(got, c.want, c.precision) {
			t.Errorf("the_most_numbers(%v) == %.1f, want %.1f", c.in, got, c.want)
		}
	}
}
