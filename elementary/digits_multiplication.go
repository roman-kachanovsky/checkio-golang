package elementary

import "strconv"

func digits_multiplication(n int) int {
	var s = strconv.Itoa(n)
	var res int = 1

	for _, c := range s {
		if string(c) != "0" {
			i, _ := strconv.Atoi(string(c))
			res *= i
		}
	}
	return res
}
