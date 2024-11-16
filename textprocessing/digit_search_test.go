package textprocessing_test

import (
	"advent2023/textprocessing"
	"fmt"
	"testing"
)

func TestGetLetterDigit(t *testing.T) {
	var tests = []struct {
		text string
		from int
		want int
	}{
		{"abcone2threexyz", 3, 1},
		{"two1nine", 0, 2},
		{"abcone2threexyz", 7, 3},
		{"xtwone3four", 7, 4},
		{"fiveasdf", 0, 5},
		{"sixasdf", 0, 6},
		{"sevenasdf", 0, 7},
		{"eightwothree", 0, 8},
		{"nineasdf", 0, 9},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("find %d in %s from %d", tt.want, tt.text, tt.from)
		t.Run(testname, func(t *testing.T) {
			result, err := textprocessing.GetLetterDigit([]rune(tt.text), tt.from)
			if err != nil {
				t.Fatalf("got %v, want %d", err, tt.want)
			}
			if result != tt.want {
				t.Fatalf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestGetLetterDigitBackward(t *testing.T) {
	var tests = []struct {
		text string
		from int
		want int
	}{
		{"abcone2threexyz", 5, 1},
		{"two1nine", 2, 2},
		{"abcone2threexyz", 11, 3},
		{"xtwone3four", 10, 4},
		{"fiveasdf", 3, 5},
		{"sixasdf", 2, 6},
		{"sevenasdf", 4, 7},
		{"eightwothree", 4, 8},
		{"nineasdf", 3, 9},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("find %d in %s from %d", tt.want, tt.text, tt.from)
		t.Run(testname, func(t *testing.T) {
			result, err := textprocessing.GetLetterDigitBackward([]rune(tt.text), tt.from)
			if err != nil {
				t.Fatalf("got %v, want %d", err, tt.want)
			}
			if result != tt.want {
				t.Fatalf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestMustGetLastDigit(t *testing.T) {
	var tests = []struct {
		text string
		want int
	}{
		{"abcone", 1},
		{"1ninetwo", 2},
		{"abcone2three", 3},
		{"abcone2three3ne", 3},
		{"xtwone3four", 4},
		{"fiveasdf", 5},
		{"asdfsix", 6},
		{"sevenasdf", 7},
		{"twothreeight", 8},
		{"nineasdf", 9},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("find %d in %s", tt.want, tt.text)
		t.Run(testname, func(t *testing.T) {
			result := textprocessing.MustGetLastDigit([]rune(tt.text))
			if result != tt.want {
				t.Fatalf("got %d, want %d", result, tt.want)
			}
		})
	}
}
