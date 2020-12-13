# vton
Go package containing implementations of encoder and decoder for encoding vowels/numbers.

| Latter  | Number  |
| ---     | ---     |
|'a'      | '1'     |
|'e'      | '2'     |
|'i'      | '3'     |
|'o'      | '4'     |
|'u'      | '5'     |

### Example

    package main

    import (
	    "fmt"

	    "github.com/igkostyuk/vton"
    )

    func main() {
	    enc := vton.NewEncoding()

	    d := "hello"
	    ec := enc.Encode(d)
	    fmt.Printf("encoded %s is %s.\n", d, ec)

	    e := "h2ll4"
	    dc := enc.Decode(e)
	    fmt.Printf("decoded %s is %s.\n", e, dc)
    }
