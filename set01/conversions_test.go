package cryptography

import (
	"testing"
)

func TestStringToBase64(t *testing.T) {
	const (
		TestFunc      = "StringToBase64"
		TestGotColor  = "got: \033[1;31m%v\033[0m\t"
		TestWantColor = "want: \033[1;33m%v\033[0m"
		ColorPrint    = TestFunc + TestGotColor + TestWantColor
	)
	var tests = []struct {
		input string
		want  string
	}{
		// {"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
		{"this is a test", "dGhpcyBpcyBhIHRlc3Q="},
	}

	for _, test := range tests {
		if got := StringToBase64(test.input); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}

func TestHexToString(t *testing.T) {
	const (
		TestFunc      = "HexToString"
		TestGotColor  = "\tgot: \033[1;31m%v\033[0m\t"
		TestWantColor = "want: \033[1;33m%v\033[0m"
		ColorPrint    = TestFunc + TestGotColor + TestWantColor
	)
	var tests = []struct {
		input string
		want  string
	}{
		{"48656c6c6f20476f7068657221", "Hello Gopher!"},
	}

	for _, test := range tests {
		if got := HexToString(test.input); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}
