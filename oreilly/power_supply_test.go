package oreilly

import (
	"testing"
	"reflect"
)

func TestPowerSupply(t *testing.T) {
	cases := []struct{
		network [][]string
		plants map[string]int
		want []string
	}{
		{
			[][]string{
				{"p1", "c1"},
				{"c1", "c2"},
			},
			map[string]int{"p1": 1},
			[]string{"c2"},
		},
		{
			[][]string{
				{"c0", "c1"},
				{"c1", "p1"},
				{"c1", "c3"},
				{"p1", "c4"},
			},
			map[string]int{"p1": 1},
			[]string{"c0", "c3"},
		},
		{
			[][]string{
				{"p1", "c1"},
				{"c1", "c2"},
				{"c2", "c3"},
			},
			map[string]int{"p1": 3},
			[]string{},
		},
		{
			[][]string{
				{"c0", "p1"},
				{"p1", "c2"},
			},
			map[string]int{"p1": 0},
			[]string{"c0", "c2"},
		},
		{
			[][]string{
				{"p0", "c1"},
				{"p0", "c2"},
				{"c2", "c3"},
				{"c3", "p4"},
				{"p4", "c5"},
			},
			map[string]int{"p0": 1, "p4": 1},
			[]string{},
		},
		{
			[][]string{
				{"c0", "p1"},
				{"p1", "c2"},
				{"c2", "c3"},
				{"c2", "c4"},
				{"c4", "c5"},
				{"c5", "c6"},
				{"c5", "p7"},
			},
			map[string]int{"p1": 1, "p7": 1},
			[]string{"c3", "c4", "c6"},
		},
		{
			[][]string{
				{"p0", "c1"},
				{"p0", "c2"},
				{"p0", "c3"},
				{"p0", "c4"},
				{"c4", "c9"},
				{"c4", "c10"},
				{"c10", "c11"},
				{"c11", "p12"},
				{"c2", "c5"},
				{"c2", "c6"},
				{"c5", "c7"},
				{"c5", "p8"},
			},
			map[string]int{"p0": 1, "p12": 4, "p8": 1},
			[]string{"c6", "c7"},
		},
		{
			[][]string{
				{"c1", "c2"},
				{"c2", "c3"},
			},
			map[string]int{},
			[]string{"c1", "c2", "c3"},
		},
		{
			[][]string{
				{"p1", "c2"},
				{"p1", "c4"},
				{"c4", "c3"},
				{"c2", "c3"},
			},
			map[string]int{"p1": 1},
			[]string{"c3"},
		},

		{
			[][]string{
				{"p1", "c2"},
				{"p1", "c4"},
				{"c2", "c3"},
			},
			map[string]int{"p1": 4},
			[]string{},
		},
	}

	for _, c := range cases {
		got := power_supply(c.network, c.plants)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("power_supply(%q, %v) == %q want %q", c.network, c.plants, got, c.want)
		}
	}
}