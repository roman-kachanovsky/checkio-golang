package codeship

import "../tools"

func golden_pyramid(pyramid [][]int) int {
	tmp := make([][]int, len(pyramid))
	copy(tmp, pyramid)

	for i := len(pyramid) - 2; i >= 0; i-- {
		for j := 0; j < i + 1; j++ {
			tmp[i][j] += tools.MaxFromPairOfInt(tmp[i + 1][j], tmp[i + 1][j + 1])
		}
	}

	return tmp[0][0]
}
