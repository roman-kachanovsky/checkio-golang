package tools

import (
	"testing"
	"reflect"
)

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
			t.Errorf("SumInt(%v) == %d, want %d", c.in, got, c.want)
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
			t.Errorf("MaxFloat64(%v) == %.2f, want %.2f", c.in, got, c.want)
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
			t.Errorf("MinFloat64(%v) == %.2f, want %.2f", c.in, got, c.want)
		}
	}
}

func TestToFixed(t *testing.T) {
	cases := []struct {
		in float64
		precision int
		want float64
	}{
		{0.000001, 1, 0.0},
		{0.5, 0, 1.0},
		{0.45, 1, 0.5},
		{0.1234, 2, 0.12},
		{0.12345, 4, 0.1235},
	}

	for _, c := range cases {
		got := ToFixed(c.in, c.precision)
		if got != c.want {
			t.Errorf("ToFixed(%.5f, %d) == %.5f, want %.5f", c.in, c.precision, got, c.want)
		}
	}
}

func TestAllBool(t *testing.T) {
	cases := []struct {
		in []bool
		want bool
	}{
		{[]bool{}, false},
		{[]bool{true}, true},
		{[]bool{false}, false},
		{[]bool{true, true ,true}, true},
		{[]bool{true, false, true}, false},
		{[]bool{true, true, false}, false},
		{[]bool{false, false, false}, false},
	}

	for _, c := range cases {
		got := AllBool(c.in)
		if got != c.want {
			t.Errorf("AllBool(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestAnyBool(t *testing.T) {
	cases := []struct {
		in []bool
		want bool
	}{
		{[]bool{}, false},
		{[]bool{true}, true},
		{[]bool{false}, false},
		{[]bool{true, true ,true}, true},
		{[]bool{true, false, true}, true},
		{[]bool{false, true, false}, true},
		{[]bool{false, false, false}, false},
	}

	for _, c := range cases {
		got := AnyBool(c.in)
		if got != c.want {
			t.Errorf("AnyBool(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestStringInSlice(t *testing.T) {
	cases := []struct {
		slice []string
		value string
		want bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{}, "a", false},
	}

	for _, c := range cases {
		got := StringInSlice(c.slice, c.value)
		if got != c.want {
			t.Errorf("StringInSlice(%q, %q) == %v, want %v", c.slice, c.value, got, c.want)
		}
	}
}

func TestIntInSlice(t *testing.T) {
	cases := []struct {
		slice []int
		value int
		want bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{}, 1, false},
	}

	for _, c := range cases {
		got := IntInSlice(c.slice, c.value)
		if got != c.want {
			t.Errorf("IntInSlice(%v, %d) == %v, want %v", c.slice, c.value, got, c.want)
		}
	}
}

func TestFloatInSlice(t *testing.T) {
	cases := []struct {
		slice []float64
		value float64
		want bool
	}{
		{[]float64{1.2, 2.5, 3.7}, 2.5, true},
		{[]float64{1.2, 2.5, 3.7}, 1.3, false},
		{[]float64{}, 2.9, false},
	}

	for _, c := range cases {
		got := FloatInSlice(c.slice, c.value)
		if got != c.want {
			t.Errorf("FloatInSlice(%v, %.1f) == %v, want %v", c.slice, c.value, got, c.want)
		}
	}
}

func TestMakeSetFromStringSlice(t *testing.T) {
	cases := []struct {
		in []string
		want []string
	}{
		{[]string{"a", "b", "a", "a", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{}, []string{}},
	}

	for _, c := range cases {
		got := MakeSetFromStringSlice(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MakeSetFromStringSlice(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMakeSetFromIntSlice(t *testing.T) {
	cases := []struct {
		in []int
		want []int
	}{
		{[]int{1, 2, 1, 1, 3, 4, 3, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
	}

	for _, c := range cases {
		got := MakeSetFromIntSlice(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MakeSetFromIntSlice(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMakeSetFromFloatSlice(t *testing.T) {
	cases := []struct {
		in []float64
		want []float64
	}{
		{[]float64{1.2, 2.0, 1.3, 1.2, 3.4, 4.1, 3.4, 5.0}, []float64{1.2, 2.0, 1.3, 3.4, 4.1, 5.0}},
		{[]float64{1.2, 1.3, 2.0, 3.4, 4.1, 5.0}, []float64{1.2, 1.3, 2.0, 3.4, 4.1, 5.0}},
		{[]float64{}, []float64{}},
	}

	for _, c := range cases {
		got := MakeSetFromFloatSlice(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MakeSetFromFloatSlice(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSubtractStringSets(t *testing.T) {
	cases := []struct {
		a []string
		b []string
		want []string
	}{
		{[]string{"a", "b", "c"}, []string{"b", "c"}, []string{"a"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{}},
		{[]string{"a", "b", "c"}, []string{"b"}, []string{"a", "c"}},
		{[]string{}, []string{"b", "c"}, []string{}},
		{[]string{"a", "b", "c"}, []string{}, []string{"a", "b", "c"}},
	}

	for _, c := range cases {
		got := SubtractStringSets(c.a, c.b)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SubtractStringSets(%q, %q) == %q, want %q", c.a, c.b, got, c.want)
		}
	}
}

func TestSubtractIntSets(t *testing.T) {
	cases := []struct {
		a []int
		b []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{2, 3}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{2}, []int{1, 3}},
		{[]int{}, []int{2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, c := range cases {
		got := SubtractIntSets(c.a, c.b)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SubtractIntSets(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestSubtractFloatSets(t *testing.T) {
	cases := []struct {
		a []float64
		b []float64
		want []float64
	}{
		{[]float64{1.1, 2.2, 3.3}, []float64{2.2, 3.3}, []float64{1.1}},
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}, []float64{}},
		{[]float64{1.1, 2.2, 3.3}, []float64{2.2}, []float64{1.1, 3.3}},
		{[]float64{}, []float64{2.2, 3.3}, []float64{}},
		{[]float64{1.1, 2.2, 3.3}, []float64{}, []float64{1.1, 2.2, 3.3}},
	}

	for _, c := range cases {
		got := SubtractFloatSets(c.a, c.b)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SubtractFloatSets(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}