package main

import "testing"

import "fmt"

const (
	TestFunc      = "ConvertHexToBase64"
	TestGotColor  = "got: \033[1;31m%v\033[0m\t"
	TestWantColor = "want: \033[1;33m%v\033[0m"
	ColorPrint    = TestFunc + TestGotColor + TestWantColor
)

// Challenge one test
func TestConvertHexToBase64(t *testing.T) {
	fmt.Println("Testing challenge one...")
	var tests = []struct {
		input string
		want  string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}

	for _, test := range tests {
		if got := ConvertHexToBase64(test.input); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}
