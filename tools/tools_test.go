package tools

import "testing"

func TestMaxFromPairOfInt(t *testing.T) {
	cases := []struct {
		a int
		b int
		want int
	}{
		{1, 2, 2},
		{1, 1, 1},
		{2, 1, 2},
		{0, -1, 0},
		{-1, 0, 0},
		{-1, -1, -1},
	}

	for _, c := range cases {
		got := MaxFromPairOfInt(c.a, c.b)
		if got != c.want {
			t.Errorf("MaxFromPairOfInt(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestMinFromPairOfInt(t *testing.T) {
	cases := []struct {
		a int
		b int
		want int
	}{
		{1, 2, 1},
		{1, 1, 1},
		{2, 1, 1},
		{0, -1, -1},
		{-1, 0, -1},
		{-1, -1, -1},
	}

	for _, c := range cases {
		got := MinFromPairOfInt(c.a, c.b)
		if got != c.want {
			t.Errorf("MinFromPairOfInt(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestSumInt(t *testing.T) {
	cases := []struct {
		in []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 2, 3}, 6},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 28},
		{[]int{5}, 5},
		{[]int{-2, -1, 0, 1, 2}, 0},
	}

	for _, c := range cases {
		got := SumInt(c.in)
		if got != c.want {
			t.Errorf("SumInt(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestMaxFloat64(t *testing.T) {
	cases := []struct {
		in []float64
		want float64
	}{
		{[]float64{}, 0.0},
		{[]float64{1, 1, 1}, 1.0},
		{[]float64{1, 2, 3}, 3.0},
		{[]float64{0, 1, 2, 9, 4, 5, 6, 7}, 9.0},
		{[]float64{5}, 5.0},
		{[]float64{-2, -1, 0, 1, 2}, 2.0},
	}

	for _, c := range cases {
		got := MaxFloat64(c.in)
		if got != c.want {
			t.Errorf("MaxFloat64(%q) == %.2f, want %.2f", c.in, got, c.want)
		}
	}
}

func TestMinFloat64(t *testing.T) {
	cases := []struct {
		in []float64
		want float64
	}{
		{[]float64{}, 0.0},
		{[]float64{1, 1, 1}, 1.0},
		{[]float64{1, 2, 3}, 1.0},
		{[]float64{0, 1, 2, 9, 4, 5, 6, 7}, 0.0},
		{[]float64{5}, 5.0},
		{[]float64{-2, -1, 0, 1, 2}, -2.0},
	}

	for _, c := range cases {
		got := MinFloat64(c.in)
		if got != c.want {
			t.Errorf("MinFloat64(%q) == %.2f, want %.2f", c.in, got, c.want)
		}
	}
}