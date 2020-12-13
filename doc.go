// Package vton containing implementations of
// encoder and decoder for encoding vowels/numbers.
//
// Usage
//
//  enc := vton.NewEncoding()
//	d := "hello"
//	ec := enc.Encode(d)
//	fmt.Printf("encoded %s is %s.\n", d, ec)
//	e := "h2ll4"
//	dc := enc.Decode(e)
//	fmt.Printf("decoded %s is %s.\n", e, dc)
//
package vton
