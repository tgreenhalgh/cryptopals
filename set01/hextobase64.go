// https://cryptopals.com/sets/1/challenges/1

package cryptography

import b64 "encoding/base64"

// HexToBase64 takes a string in hex and returns a base64 string
func HexToBase64(data string) string {

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
	// fmt.Println(sEnc)

	// sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	// fmt.Println(string(sDec))
	// fmt.Println()

	// uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	// fmt.Println(uEnc)
	// uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	// fmt.Println(string(uDec))
}
