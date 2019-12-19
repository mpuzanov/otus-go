package topwords

import (
	"strings"
	"testing"
)

func TestTopWords(t *testing.T) {
	var tests = []struct {
		input string
		wont  string
	}{
		{`111 222 111 222 111 333`, "111 222 333"},
		{`cat and dog one dog two cats and one man`, "and dog one cat cats man two"},
		{`a b c d b`, "b a c d"},
	}
	for _, test := range tests {
		if got := strings.Join(TopWords10(test.input), " "); got != test.wont {
			t.Errorf("convert(%q) = %q, wont=%q", test.input, got, test.wont)
		}
	}
}
