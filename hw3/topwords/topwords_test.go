package topwords

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTopWords(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		wont  string
	}{
		{"test 1", "111 222 111 222 111 333", "111 222 333"},
		{"test 2", "cat and dog one dog two cats and one man", "and dog one cat cats man two"},
		{"test 3", "a b c d b", "b a c d"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// if got := strings.Join(TopWords10(tC.input), " "); got != tC.wont {
			// 	 t.Errorf("convert(%q) = %q, wont=%q", tC.input, got, tC.wont)
			//  }
			assert.Equal(t, tC.wont, strings.Join(TopWords10(tC.input), " "))
		})
	}
}
