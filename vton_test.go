package vton_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/igkostyuk/vton"
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

func TestRandWords(t *testing.T) {
	e := vton.NewEncoding()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		want := RandStringRunes(25)
		got := e.Decode(e.Encode(want))
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] // #nosec
	}

	return string(b)
}

func BenchmarkEncoding(b *testing.B) {
	e := vton.NewEncoding()
	for n := 0; n < b.N; n++ {
		e.Encode("hello")
	}
}

func BenchmarkDecoding(b *testing.B) {
	e := vton.NewEncoding()
	for n := 0; n < b.N; n++ {
		e.Decode("h2ll4")
	}
}
