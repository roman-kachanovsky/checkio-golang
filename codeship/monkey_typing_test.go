package codeship

import "testing"

func TestMonkeyTyping(t *testing.T) {
	cases := []struct {
		text string
		words []string
		want int
	}{
		{
			"How aresjfhdskfhskd you?",
			[]string{"how", "are", "you", "hello"},
			3,
		},
		{
			"Bananas, give me bananas!!!",
			[]string{"banana", "bananas"},
			2,
		},
		{
			"Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
			[]string{"sum", "hamlet", "infinity", "anything"},
			1,
		},
	}

	for _, c := range cases {
		got := monkey_typing(c.text, c.words)
		if got != c.want {
			t.Errorf("monkey_typing(%q, %q) == %d, want %d", c.text, c.words, got, c.want)
		}
	}
}


