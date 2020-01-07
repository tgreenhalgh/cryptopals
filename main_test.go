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

// Challenge two test
func TestExerciseTwo(t *testing.T) {
	fmt.Println("Testing challenge two...")
	const (
		TestFunc      = "exerciseTwo"
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
		if got := exerciseTwo(test.input1, test.input2); got != test.want {
			t.Errorf(ColorPrint, got, test.want)
		}
	}
}
