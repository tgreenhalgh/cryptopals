package cryptography

// from https://en.wikipedia.org/wiki/Letter_frequency
// and http://www.data-compression.com/english.html for " " (spaces)
var englishFreq = map[string]float32{
	"a": .08167,
	"b": .01492,
	"c": .02782,
	"d": .04253,
	"e": .12702,
	"f": .02228,
	"g": .02015,
	"h": .06094,
	"i": .06966,
	"j": .00153,
	"k": .00772,
	"l": .04025,
	"m": .02406,
	"n": .06749,
	"o": .07507,
	"p": .01929,
	"q": .00095,
	"r": .05987,
	"s": .06327,
	"t": .09056,
	"u": .02758,
	"v": .00978,
	"w": .02360,
	"x": .00150,
	"y": .01974,
	"z": .00074,
	" ": .19182,
}

// CalcFreq is a test function
func CalcFreq(s string) float32 {
	var score float32

	for i := 0; i < len(s); i++ {
		score += englishFreq[string(s[i])]
	}

	return score
}
