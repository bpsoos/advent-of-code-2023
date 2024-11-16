package day2_test

import (
	"advent2023/day2"
	"reflect"
	"strings"
	"testing"
)

func TestParseGame(t *testing.T) {
	rawGames := strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        `)
	want := []day2.Game{
		{
			Id: 1,
			RevealedSets: []day2.CubeSet{
				{Blue: 3, Red: 4},
				{Red: 1, Green: 2, Blue: 6},
				{Green: 2},
			},
		},
		{
			Id: 2,
			RevealedSets: []day2.CubeSet{
				{Blue: 1, Green: 2},
				{Red: 1, Green: 3, Blue: 4},
				{Green: 1, Blue: 1},
			},
		},
	}
	games := day2.ParseGames(rawGames)
	if len(want) != len(games) {
		t.Errorf("lenghts don't match")
	}
	if !reflect.DeepEqual(games, want) {
		t.Errorf("parsed games don't match expected games")
		t.Logf("parsed games: %v", games)
		t.Logf("expected games: %v", want)
	}
}

func TestParseGameWithTwoDigitNumbers(t *testing.T) {
	rawGames := strings.NewReader(`Game 11: 33 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 22: 1 blue, 2 green; 3 green, 4 blue, 12 red; 11 green, 1 blue
        `)
	want := []day2.Game{
		{
			Id: 11,
			RevealedSets: []day2.CubeSet{
				{Blue: 33, Red: 4},
				{Red: 1, Green: 2, Blue: 6},
				{Green: 2},
			},
		},
		{
			Id: 22,
			RevealedSets: []day2.CubeSet{
				{Blue: 1, Green: 2},
				{Red: 12, Green: 3, Blue: 4},
				{Green: 11, Blue: 1},
			},
		},
	}
	games := day2.ParseGames(rawGames)
	if len(want) != len(games) {
		t.Errorf("lenghts don't match")
	}
	if !reflect.DeepEqual(games, want) {
		t.Errorf("parsed games don't match expected games")
		t.Logf("parsed games: %v", games)
		t.Logf("expected games: %v", want)
	}
}
