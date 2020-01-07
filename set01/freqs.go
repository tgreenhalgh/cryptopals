package cryptography

var englishFreq = map[string]float32{
	"a": .8167,
	"b": .1492,
	"c": .2782,
	"d": .4253,
	"e": .12702,
	"f": .2228,
	"g": .2015,
	"h": .6094,
	"i": .6966,
	"j": .0153,
	"k": .0772,
	"l": .4025,
	"m": .2406,
	"n": .6749,
	"o": .7507,
	"p": .1929,
	"q": .0095,
	"r": .5987,
	"s": .6327,
	"t": .9056,
	"u": .2758,
	"v": .0978,
	"w": .2360,
	"x": .0150,
	"y": .1974,
	"z": .074,
}

// CalcFreq is a test function
func CalcFreq(s string) float32 {
	var score float32

	for i := 0; i < len(s); i++ {
		score += englishFreq[string(s[i])]
	}

	return score
}
