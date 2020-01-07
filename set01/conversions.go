// https://cryptopals.com/sets/1/challenges/1

package cryptography

import (
	b64 "encoding/base64"
	hex "encoding/hex"
	"fmt"
	"log"
	"strconv"
)

// WARNING:
// Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

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

// HexToBin converts a hex string to bin
func HexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}

	// %016b indicates base 2, zero padded, with 16 characters
	return fmt.Sprintf("%016b", ui), nil
}

// FixedXOR takes two stringd of HEX and XORs them, returning a HEX string
func FixedXOR(s1 string, s2 string) string {
	if len(s1) != len(s2) {
		panic("strings don't have equal lengths")
	}
	var hex = ""
	for i := 0; i < len(s1); i += 2 {
		h1, err := strconv.ParseInt("0x"+string(s1[i:i+2]), 0, 64)
		if err != nil {
			panic(err)
		}
		h2, err := strconv.ParseInt("0x"+string(s2[i:i+2]), 0, 64)
		if err != nil {
			panic(err)
		}
		temp := strconv.FormatInt(h1^h2, 16)
		if err != nil {
			panic(err)
		}
		hex += temp
	}
	return hex
}
