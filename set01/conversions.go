// https://cryptopals.com/sets/1/challenges/1

package cryptography

import (
	b64 "encoding/base64"
	hex "encoding/hex"
	"log"
)

// StringToBase64 takes a string in hex and returns a base64 string
func StringToBase64(data string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}

// HexToASCIIString takes a string in hex and returns []byte
func HexToASCIIString(s string) string {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return string(decoded)
}
