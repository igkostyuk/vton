package vton

import (
	"strings"
	"unicode"
)

type Encoding struct {
	letters map[rune]rune
	numbers map[rune]rune
}

func NewEncoding() Encoding {
	vtnencoding := []struct {
		letter rune
		number rune
	}{
		{letter: 'a', number: '1'},
		{letter: 'e', number: '2'},
		{letter: 'i', number: '3'},
		{letter: 'o', number: '4'},
		{letter: 'u', number: '5'},
	}
	var e Encoding
	e.letters = make(map[rune]rune, len(vtnencoding))
	e.numbers = make(map[rune]rune, len(vtnencoding))
	for _, c := range vtnencoding {
		e.numbers[c.number] = c.letter
		e.letters[c.letter] = c.number
		e.letters[unicode.ToUpper(c.letter)] = c.number
	}

	return e
}

func (e Encoding) Decode(str string) string {
	return strings.Map(findRune(e.numbers), str)
}

func (e Encoding) Encode(str string) string {
	return strings.Map(findRune(e.letters), str)
}

func findRune(m map[rune]rune) func(rune) rune {
	return func(r rune) rune {
		if v, ok := m[r]; ok {
			return v
		}

		return r
	}
}
