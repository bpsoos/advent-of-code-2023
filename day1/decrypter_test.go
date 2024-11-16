package day1_test

import (
	"advent2023/day1"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestDecrypter(t *testing.T) {
	var tests = []struct {
		calibrationDocument io.Reader
		want                int
	}{
		{strings.NewReader(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`), 142},
	}
	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := day1.Decrypter{}.Decrypt(tt.calibrationDocument)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}
