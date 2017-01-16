package mine

import "../tools"

func p(white, black, step int) float64 {
	var pWhite = float64(white) / float64(white + black)
	var pBlack = 1.0 - pWhite

	if step == 1 {
		return pWhite
	} else {
		return p(white - 1, black + 1, step - 1) * pWhite + p(white + 1, black - 1, step - 1) * pBlack
	}
}

func pearls_in_the_box(s string, n int) float64 {
	var bCount, wCount int

	for _, c := range s {
		if c == 'w' {
			wCount++
		} else {
			bCount++
		}
	}

	return tools.ToFixed(p(wCount, bCount, n), 2)
}
