package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	var tests = []struct {
		input string
		wont  string
		err   error
	}{
		{`a4bc2d5e`, `aaaabccddddde`, nil},
		{`abcd`, `abcd`, nil},
		{`п5г4`, `пппппгггг`, nil},
		{`45`, ``, errBadString},
		{``, ``, nil},
		{`qwe\4\5`, `qwe45`, nil},
		{`qwe\45`, `qwe44444`, nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}
	for _, test := range tests {
		if got, err := Unpack(test.input); got != test.wont || err != test.err {
			t.Errorf("convert(%v) = %v, wont=%v, error=%v", test.input, got, test.wont, err)
		}
	}
}
