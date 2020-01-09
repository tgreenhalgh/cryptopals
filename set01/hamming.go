package cryptography

import (
	"fmt"
)

// converts two strings to binary, then checks the Hamming Distance
// https://sciencing.com/how-to-calculate-hamming-distance-12751770.html

// XORBytes takes two byte slices and XORs them together, returning the final
// byte slice. It is an error to pass in two byte slices that do not have the
// same length.
func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of byte slices is not equivalent: %d != %d", len(a), len(b))
	}

	buf := make([]byte, len(a))

	for i := range a {
		buf[i] = a[i] ^ b[i]
	}

	return buf, nil
}

func stringToBin(s string) string {
	str := ""
	for _, c := range s {
		str = fmt.Sprintf("%s%.8b", str, c)
	}
	return str
}

// HammingDist converts two equal length strings to binary
// then checks the difference in bits.
// returns -1 if strings are not equal length
func HammingDist(s1, s2 string) int {
	var diff int

	if len(s1) != len(s2) {
		return -1
	}

	bin1 := stringToBin(s1)
	bin2 := stringToBin(s2)

	res, err := XORBytes([]byte(bin1), []byte(bin2))
	if err != nil {
		panic(err)
	}

	for _, v := range res {
		diff += int(v)
	}

	return diff
}
