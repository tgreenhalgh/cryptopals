package main

import (
	cr "cryptopals/set01"
	"fmt"
	"math"
	"strconv"
	"strings"
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

/*
	Exercise 3
*/
// Single-byte XOR cipher
func exerciseThree(s1 string) string {
	// expected is from English freq tables, looking at "ETAOIN"
	// const expected = .51147
	// expected is from English freq tables, looking at "ETAOIN SHRDLU"
	const expected = 0.99773
	scoreDiff := 1.0
	var cleartext string
	var hexChar string

	// XOR by a single byte
	for i := 0; i < 256; i++ {
		// convert to hex
		hexChar = strconv.FormatInt(int64(i), 16)
		if len(hexChar)%2 != 0 {
			hexChar = "0" + hexChar
		}

		repString := strings.Repeat(hexChar, len(s1)/2)
		hex := cr.FixedXOR(s1, repString)
		text := cr.HexToASCIIString(hex)
		tempScore := cr.DecodedFreq(text)

		// looking for the closest distribution to English
		if math.Abs(expected-tempScore) < scoreDiff {
			scoreDiff = math.Abs(expected - tempScore)
			cleartext = text
		}
	}

	return cleartext
}

func main() {
	fmt.Println("Exercise one:\n", ConvertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println("Exercise two:\n", exerciseTwo("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
	fmt.Println("Exercise three:\n", exerciseThree("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
}
