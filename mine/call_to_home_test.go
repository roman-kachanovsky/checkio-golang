package mine

import "testing"

func TestCallToHome(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{
			"2014-01-01 01:12:13 181",
			"2014-01-02 20:11:10 600",
			"2014-01-03 01:12:13 6009",
			"2014-01-03 12:13:55 200",
		}, 124},
		{[]string{
			"2014-02-05 01:00:00 1",
			"2014-02-05 02:00:00 1",
			"2014-02-05 03:00:00 1",
			"2014-02-05 04:00:00 1",
		}, 4},
		{[]string{
			"2014-02-05 01:00:00 60",
			"2014-02-05 02:00:00 60",
			"2014-02-05 03:00:00 60",
			"2014-02-05 04:00:00 6000",
		}, 106},
	}

	for _, c := range cases {
		got := call_to_home(c.in)
		if got != c.want {
			t.Errorf("call_to_home(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
