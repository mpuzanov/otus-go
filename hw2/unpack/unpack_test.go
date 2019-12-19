package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		wont  string
		err   error
	}{
		{"test1", `a4bc2d5e`, `aaaabccddddde`, nil},
		{"test2", `abcd`, `abcd`, nil},
		{"test3", `п5г4`, `пппппгггг`, nil},
		{"test4", `45`, ``, errBadString},
		{"test5", ``, ``, nil},
		{"test6", `qwe\4\5`, `qwe45`, nil},
		{"test7", `qwe\45`, `qwe44444`, nil},
		{"test8", `qwe\\5`, `qwe\\\\\`, nil},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got, err := Unpack(tC.input); got != tC.wont || err != tC.err {
				t.Errorf("convert(%v) = %v, wont=%v, error=%v", tC.input, got, tC.wont, err)
			}
		})
	}
}

func BenchmarkUnpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unpack(`a4bc2d5e`)
	}
}

func BenchmarkUnpack2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unpack2(`a4bc2d5e`)
	}
}
