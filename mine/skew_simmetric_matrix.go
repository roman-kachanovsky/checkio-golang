package mine

import "../tools"

func skew_simmetric_matrix(m [][]int) bool {
	var a []bool

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			a = append(a, m[i][j] == -m[j][i])
		}
	}

	return tools.AllBool(a)
}