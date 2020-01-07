package main

import (
	cr "cryptopals/set01"
	"fmt"
)

// ConvertHexToBase64 takes a string of hex, converts it to ascii, then converts that to base64
func ConvertHexToBase64(s string) string {
	hex := cr.HexToAsciitring(s)
	b := cr.StringToBase64(hex)
	return b
}

func main() {
	fmt.Println("Code for challenge one")
}
