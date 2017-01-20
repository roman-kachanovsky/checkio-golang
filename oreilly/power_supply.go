package oreilly

func stringInSlice(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func findShortestPath(g map[string][]string, s, e string, p []string) []string {
	p = append(p, s)

	if s == e {
		return p
	}

	if _, ok := g[s]; !ok {
		return nil
	}

	var shortest []string = nil

	for _, n := range g[s] {
		if !stringInSlice(p, n) {
			new_path := findShortestPath(g, n, e, p)
			if new_path != nil && (shortest == nil || len(new_path) < len(shortest)) {
				shortest = new_path
			}
		}
	}

	return shortest
}

func removeDoublets(a []string) []string {
	clean_slice := []string{}

	for _, v := range a {
		if !stringInSlice(clean_slice, v) {
			clean_slice = append(clean_slice, v)
		}
	}

	return clean_slice
}

func subtractSlices(a, b []string) []string {
	res := []string{}

	for _, v := range a {
		if !stringInSlice(b, v) {
			res = append(res, v)
		}
	}

	return res
}

func power_supply(network [][]string, plants map[string]int) []string {
	var all_points, cities, plant_names, covered_cities []string

	for _, e := range network {
		for _, v := range e {
			if v[0] == 99 {
				cities = append(cities, v)
			} else {
				plant_names = append(plant_names, v)
			}

			if !stringInSlice(all_points, v) {
				all_points = append(all_points, v)
			}
		}
	}

	cities = removeDoublets(cities)
	supply_graph := make(map[string][]string, len(all_points))

	for _, k := range all_points {
		for _, m := range network {
			if stringInSlice(m, k) {
				if k != m[0] {
					supply_graph[k] = append(supply_graph[k], m[0])
				} else {
					supply_graph[k] = append(supply_graph[k], m[1])
				}
			}
		}
	}

	for plant, supply_range := range plants {
		for _, city := range cities {
			path := findShortestPath(supply_graph, city, plant, []string{})

			if path != nil && len(path) - 1 <= supply_range {
				covered_cities = append(covered_cities, city)
			}
		}
	}

	covered_cities = removeDoublets(covered_cities)

	return subtractSlices(cities, covered_cities)
}
