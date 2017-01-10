package elementary

import (
	"sort"
	"math"
)

type ByAbs []int

func (a ByAbs) Len() int { return len(a) }
func (a ByAbs) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAbs) Less(i, j int) bool { return math.Abs(float64(a[i])) < math.Abs(float64(a[j])) }

func absolute_sorting(a []int) []int {
	sort.Sort(ByAbs(a))
	return a
}
