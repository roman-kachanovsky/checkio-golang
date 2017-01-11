package elementary

import "testing"

func TestRightToLeft(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{[]string{"left", "right", "left", "stop"}, "left,left,left,stop"},
		{[]string{"bright aright", "ok"}, "bleft aleft,ok"},
		{[]string{"brightness wright"}, "bleftness wleft"},
		{[]string{"enough", "jokes"}, "enough,jokes"},
	}

	for _, c := range cases {
		got := right_to_left(c.in)
		if got != c.want {
			t.Errorf("right_to_left(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

