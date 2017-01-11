package elementary

import "testing"

func TestSecretMessage(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"How are you? Eh, ok. Low or Lower? Ohhh.", "HELLO"},
		{"hello world!", ""},
		{"HELLO WORLD!!!", "HELLOWORLD"},
	}

	for _, c := range cases {
		got := secret_message(c.in)
		if got != c.want {
			t.Errorf("secret_message(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

