// https://cryptopals.com/sets/1/challenges/1

package cryptography

import (
	hex "encoding/hex"
	"strconv"
)

/*
Encrypt it, under the key "ICE", using repeating-key XOR.

In repeating-key XOR, you'll sequentially apply each byte
of the key; the first byte of plaintext will be XOR'd
against I, the next C, the next E, then I again for the
4th byte, and so on.
*/

// RKXOR is a repeating key XOR
// pass key as a string, plaintext as string, returns string
func RKXOR(key string, pt string) string {
	keyLength := len(key) * 2
	var xord string

	src := []byte(pt)
	encodedPt := hex.EncodeToString(src)

	src = []byte(key)
	encodedKey := hex.EncodeToString(src)

	for i := 0; i < len(encodedPt); i += 2 {
		h1, err := strconv.ParseInt("0x"+string(encodedPt[i:i+2]), 0, 64)
		if err != nil {
			panic(err)
		}

		enc := string(encodedKey[i%keyLength : i%keyLength+2])
		h2, err := strconv.ParseInt("0x"+enc, 0, 64)
		if err != nil {
			panic(err)
		}

		temp := strconv.FormatInt(h1^h2, 16)
		if len(temp) == 1 {
			temp = "0" + temp
		}
		xord += temp
	}

	return xord
}
