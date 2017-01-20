package tools

import "math"

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

func round(n float64) int {
	return int(n + math.Copysign(0.5, n))
}

func ToFixed(n float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(n * output)) / output
}

func AllBool(a []bool) bool {
	if len(a) == 0 {
		return false
	}

	for _, v := range a {
		if !v {
			return false
		}
	}
	return true
}

func AnyBool(a []bool) bool {
	if len(a) == 0 {
		return false
	}

	for _, v := range a {
		if v {
			return true
		}
	}
	return false
}

func StringInSlice(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IntInSlice(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func FloatInSlice(slice []float64, value float64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func MakeSetFromStringSlice(slice []string) []string {
	set := []string{}
	for _, v := range slice {
		if !StringInSlice(set, v) {
			set = append(set, v)
		}
	}
	return set
}

func MakeSetFromIntSlice(slice []int) []int {
	set := []int{}
	for _, v := range slice {
		if !IntInSlice(set, v) {
			set = append(set, v)
		}
	}
	return set
}

func MakeSetFromFloatSlice(slice []float64) []float64 {
	set := []float64{}
	for _, v := range slice {
		if !FloatInSlice(set, v) {
			set = append(set, v)
		}
	}
	return set
}

func SubtractStringSets(a, b []string) []string {
	result := []string{}
	for _, v := range a {
		if !StringInSlice(b, v) {
			result = append(result, v)
		}
	}
	return result
}

func SubtractIntSets(a, b []int) []int {
	result := []int{}
	for _, v := range a {
		if !IntInSlice(b, v) {
			result = append(result, v)
		}
	}
	return result
}

func SubtractFloatSets(a, b []float64) []float64 {
	result := []float64{}
	for _, v := range a {
		if !FloatInSlice(b, v) {
			result = append(result, v)
		}
	}
	return result
}

// TODO: Add check if slice is set
// TODO: Implement union, intersection and other set operations
