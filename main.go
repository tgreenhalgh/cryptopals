package main

import (
	cr "cryptopals/set01"
	"fmt"
)

/*
	Exercise 1
*/

// ConvertHexToBase64 takes a string of hex, converts it to ascii, then converts that to base64
func ConvertHexToBase64(s string) string {
	hex := cr.HexToASCIIString(s)
	b := cr.StringToBase64(hex)
	return b
}

/*
	Exercise 2
*/
func exerciseTwo(s1 string, s2 string) string {
	return cr.FixedXOR(s1, s2)
}

func main() {
	fmt.Println("Exercise one:\n", ConvertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println("Exercise two:\n", exerciseTwo("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
}
