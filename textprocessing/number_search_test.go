package textprocessing_test

import (
	"advent2023/textprocessing"
	"testing"
)

func TestMustGetNumberWithNumberAtEnd(t *testing.T) {
	number, length := textprocessing.MustGetNumber([]rune("...76"), 3)
	if number != 76 {
		t.Errorf("number doesn't match expected: %v", number)
	}
	if length != 2 {
		t.Errorf("number doesn't match expected: %v", length)
	}
}

func TestMustGetNumberWithNumberAtStart(t *testing.T) {
	number, length := textprocessing.MustGetNumber([]rune("76..."), 0)
	if number != 76 {
		t.Errorf("number doesn't match expected: %v", number)
	}
	if length != 2 {
		t.Errorf("length doesn't match expected: %v", length)
	}
}
