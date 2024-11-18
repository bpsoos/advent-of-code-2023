package day4_test

import (
	"advent2023/day4"
	"reflect"
	"strings"
	"testing"
)

func TestParseCardsWithSingleCard(t *testing.T) {
	want := day4.Scratchcards{{Id: 2, WinningNumbers: []int{1, 2, 3}, Numbers: []int{2, 0, 4}}}
	rawCards := strings.NewReader(strings.TrimSpace(`
Card  2:  1  2  3 |  2  0  4
`))
	parsedCards := day4.ParseCards(rawCards)
	if !reflect.DeepEqual(parsedCards, want) {
		t.Errorf(
			"parsed cards don't match expected \nparsed:  %v\nexpected: %v",
			parsedCards,
			want,
		)
	}
}

func TestParseCardsWithMultipleCards(t *testing.T) {
	want := day4.Scratchcards{
		{Id: 2, WinningNumbers: []int{1, 2, 3}, Numbers: []int{2, 0, 4}},
		{Id: 33, WinningNumbers: []int{11, 32, 0}, Numbers: []int{3, 99, 1}},
	}
	rawCards := strings.NewReader(strings.TrimSpace(`
Card  2:  1  2  3 |  2  0  4
Card 33: 11 32  0 |  3 99  1
`))
	parsedCards := day4.ParseCards(rawCards)
	if !reflect.DeepEqual(parsedCards, want) {
		t.Errorf(
			"parsed cards don't match expected \nparsed:  %v\nexpected: %v",
			parsedCards,
			want,
		)
	}
}
