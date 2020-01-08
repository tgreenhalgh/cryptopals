package cryptography

import (
	"strings"
)

// from https://en.wikipedia.org/wiki/Letter_frequency
// and http://www.data-compression.com/english.html for " " (spaces)
var englishFreq = map[string]float64{
	"e": .12702, "t": .09056, "a": .08167,
	"o": .07507, "i": .06966, "n": .06749,
	"s": .06327, "h": .06094, "r": .05987,
	"d": .04253, "l": .04025, "c": .02782,
	"u": .02758, "m": .02406, "w": .02360,
	"f": .02228, "g": .02015, "y": .01974,
	"p": .01929, "b": .01492, "v": .00978,
	"k": .00772, "j": .00153, "x": .00150,
	"q": .00095, "z": .00074, " ": .19182,
}

// DecodedFreq generates a map of the decoded string
func DecodedFreq(s string) float64 {
	var length float64
	var chi2 float64
	var freqMap = map[string]float64{}

	for i := 0; i < len(s); i++ {
		l := strings.ToLower(string(s[i]))
		if l == " " || (l >= "a" && l <= "z") {
			freqMap[string(s[i])]++
			length++
		}
	}

	for k, v := range freqMap {
		expected := length * englishFreq[strings.ToLower(k)]
		diff := v - expected
		chi2 += diff * diff / expected
	}

	return chi2
}
