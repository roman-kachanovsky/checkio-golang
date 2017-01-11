package mine

import (
	"strings"
	"strconv"
	"../tools"
)

func call_to_home(calls []string) int {
	var days = make(map[string]int)
	var durations []int

	for _, call := range calls {
		parts := strings.Split(call, " ")
		duration, _ := strconv.ParseInt(parts[2], 10, 32)
		days[parts[0]] += int((duration + 59) / 60)
	}

	for _, v := range days {
		durations = append(durations, tools.MaxFromPairOfInt(v, v * 2 - 100))
	}

	return tools.SumInt(durations)
}
