package vton_test

import (
	"goo/vton"
	"testing"
)

func TestWords(t *testing.T) {
	tt := []struct {
		decoded string
		encoded string
	}{
		{decoded: "hello", encoded: "h2ll4"},
		{decoded: "o, hi there", encoded: "4, h3 th2r2"},
		{decoded: "are you there?", encoded: "1r2 y45 th2r2?"},
	}
	e := vton.NewEncoding()
	for _, test := range tt {
		gotDecoded := e.Decode(test.encoded)
		gotEncoded := e.Encode(test.decoded)
		if gotDecoded != test.decoded {
			t.Errorf("decode got %s want %s", gotDecoded, test.decoded)
		}
		if gotEncoded != test.encoded {
			t.Errorf("encode got %s want %s", gotEncoded, test.encoded)
		}
	}
}

func TestWordsWithUpperCases(t *testing.T) {
	tt := []struct {
		decoded string
		encoded string
	}{
		{decoded: "Oyy!", encoded: "4yy!"},
		{decoded: "I will be thEre", encoded: "3 w3ll b2 th2r2"},
		{decoded: "Are you OK?", encoded: "1r2 y45 4K?"},
		{decoded: "Upper case", encoded: "5pp2r c1s2"},
	}
	e := vton.NewEncoding()
	for _, test := range tt {
		gotEncoded := e.Encode(test.decoded)
		if gotEncoded != test.encoded {
			t.Errorf("encode got %s want %s", gotEncoded, test.encoded)
		}
	}
}
