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
		{"this is a test", "dGhpcyBpcyBhIHRlc3Q="},
	}

	for _, test := range tests {
		if got := StringToBase64(test.input); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}

func TestHexToASCIIString(t *testing.T) {
	const (
		TestFunc      = "HexToASCIIString"
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
		if got := HexToASCIIString(test.input); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}

func TestFixedXOR(t *testing.T) {
	const (
		TestFunc      = "FixedXOR"
		TestGotColor  = "\tgot: \033[1;31m%v\033[0m\t"
		TestWantColor = "want: \033[1;33m%v\033[0m"
		ColorPrint    = TestFunc + TestGotColor + TestWantColor
	)
	var tests = []struct {
		input1 string
		input2 string
		want   string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}

	for _, test := range tests {
		if got := FixedXOR(test.input1, test.input2); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}
