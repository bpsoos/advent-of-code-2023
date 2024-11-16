package day1_test

import (
	"advent2023/day1"
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {
	document := strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)
	want := 142
	result := day1.Decrypter{}.Decrypt(document)
	if result != want {
		t.Errorf("got %d, want %d", result, want)
	}
}

func TestDecryptWithLetters(t *testing.T) {
	document := strings.NewReader(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)
	want := 281
	result := day1.Decrypter{}.DecryptWithLetters(document)
	if result != want {
		t.Errorf("got %d, want %d", result, want)
	}
}
