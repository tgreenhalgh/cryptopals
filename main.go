package main

import (
	"bufio"
	cr "cryptopals/set01"
	"fmt"
	"log"
	"os"
	"sort"
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
// Write a function that takes two equal-length buffers and produces their XOR combination.
func exerciseTwo(s1 string, s2 string) string {
	return cr.FixedXOR(s1, s2)
}

/*
	Exercise 3
*/
// Single-byte XOR cipher
func exerciseThree(s1 string) []string {
	var hexChar string
	var scores = []score{}
	var results []string

	// XOR by a single byte
	for i := 16; i < 256; i++ { // i = 21
		// convert to hex
		hexChar = strconv.FormatInt(int64(i), 16)
		if len(hexChar)%2 != 0 {
			hexChar = "0" + hexChar
		}

		repString := strings.Repeat(hexChar, len(s1)/2)
		hex := cr.FixedXOR(s1, repString)
		text := cr.HexToASCIIString(hex)
		tempScore := cr.DecodedFreq(text)
		if tempScore > 0 {
			item := score{id: i, score: tempScore, text: text}
			scores = append(scores, item)
		}
	}
	sort.Sort(ByScore(scores))
	// for _, t := range scores {
	// 	check := true
	// 	for _, l := range t.text {
	// 		if check && l < 19 || l > 164 {
	// 			// if check && l > 255 {
	// 			check = false
	// 			break
	// 		}
	// 	}
	// 	if check {
	// 		results = append(results, t.text)
	// 	}
	// }
	// if len(results) > 6 {
	// 	return results[:6]
	// }
	for i := 0; i < 3; i++ {
		results = append(results, scores[i].text)
	}
	return results
}

/*
Exercise four
*/
//One of the 60-character strings in the file has been encrypted by single-character XOR.
func exerciseFour() {
	decoded := [][]string{}

	file, err := os.Open("set01/04data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := exerciseThree(scanner.Text())
		if len(res) > 0 {
			decoded = append(decoded, res)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range decoded {
		for _, t := range v {
			fmt.Println(t)
		}
	}
}

type score struct {
	id    int
	score float64
	text  string
}

// ByScore implements sort.Interface for []score based on
// the score field.
type ByScore []score

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score < a[j].score }

func main() {
	// fmt.Println("Exercise one:\n", ConvertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	// fmt.Println("Exercise two:\n", exerciseTwo("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
	res := exerciseThree("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println("Exercise three:")
	for _, v := range res {
		fmt.Println(v)
	}
	// fmt.Println("Exercise four:")
	// exerciseFour()
}
