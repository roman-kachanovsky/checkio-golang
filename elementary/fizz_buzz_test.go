package elementary

import "testing"

func TestMySolution(t *testing.T) {
	cases := []struct {
		in int
		want string
	}{
		{15, "Fizz Buzz"},
		{6, "Fizz"},
		{5, "Buzz"},
		{7, "7"},
		{0, "Fizz Buzz"},
		{-1, "-1"},
	}

	for _, c := range cases {
		got := my_solution(c.in)
		if got != c.want {
			t.Errorf("my_solution(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}