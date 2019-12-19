package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var errBadString = errors.New("некорректная строка")

// Unpack Распаковка строки
func Unpack(in string) (string, error) {

	var prevChar, prevChar2 rune
	var b strings.Builder
	var out string
	var err error

	for _, val := range in {

		if unicode.IsDigit(val) {

			if prevChar == 92 && prevChar2 != 92 { // 92 - код escape - последовательности `\`
				prevChar = val
				b.WriteRune(val)
			} else {
				if num, err := strconv.Atoi(string(val)); prevChar != 0 && err == nil {
					b.WriteString(strings.Repeat(string(prevChar), num-1))
				}
			}
		}
		if prevChar == 92 && val == 92 {
			b.WriteRune(val)
		}
		if !unicode.IsDigit(val) && (prevChar != 92 && val != 92) {
			b.WriteRune(val)
		}
		if !unicode.IsDigit(val) {
			prevChar2 = prevChar
			prevChar = val
		}
	}

	out = b.String()
	if len(in) > 0 && len(out) == 0 {
		err = errBadString
	}
	return out, err
}

// Unpack2 Распаковка строки
func Unpack2(in string) (string, error) {

	var prevChar, prevChar2 rune
	var out string
	var err error

	for _, val := range in {

		if unicode.IsDigit(val) {

			if prevChar == 92 && prevChar2 != 92 { // 92 - код escape - последовательности `\`
				prevChar = val
				out += string(val)
			} else {
				if num, err := strconv.Atoi(string(val)); prevChar != 0 && err == nil {
					out += strings.Repeat(string(prevChar), num-1)
				}
			}
		}
		if prevChar == 92 && val == 92 {
			out += string(val)
		}
		if !unicode.IsDigit(val) && (prevChar != 92 && val != 92) {
			out += string(val)
		}
		if !unicode.IsDigit(val) {
			prevChar2 = prevChar
			prevChar = val
		}
	}

	if len(in) > 0 && len(out) == 0 {
		err = errBadString
	}
	return out, err
}
