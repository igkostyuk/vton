package vton

import (
	"strings"
	"unicode"
)

type encoding struct {
	latters map[rune]rune
	numbers map[rune]rune
}

func NewEncoding() encoding {
	var vtnencoding = []struct {
		latter rune
		number rune
	}{
		{latter: 'a', number: '1'},
		{latter: 'e', number: '2'},
		{latter: 'i', number: '3'},
		{latter: 'o', number: '4'},
		{latter: 'u', number: '5'},
	}
	var e encoding
	e.latters = make(map[rune]rune, len(vtnencoding))
	e.numbers = make(map[rune]rune, len(vtnencoding))
	for _, c := range vtnencoding {
		e.numbers[c.number] = c.latter
		e.latters[c.latter] = c.number
		e.latters[unicode.ToUpper(c.latter)] = c.number
	}
	return e
}

func (e encoding) Decode(str string) string {
	return strings.Map(findRune(e.numbers), str)
}

func (e encoding) Encode(str string) string {
	return strings.Map(findRune(e.latters), str)
}

func findRune(m map[rune]rune) func(rune) rune {
	return func(r rune) rune {
		if v, ok := m[r]; ok {
			return v
		}
		return r
	}
}
