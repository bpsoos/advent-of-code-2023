package day2_test

import (
	"advent2023/day2"
	"reflect"
	"testing"
)

func TestVerifyGame(t *testing.T) {
	games := []day2.Game{
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
	want := []int{2}
	result := day2.GameVerifier{
		BagContent: day2.CubeSet{Blue: 5, Green: 5, Red: 5},
	}.Verify(games)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("parsed games don't match expected games")
		t.Logf("result: %v", result)
		t.Logf("expected: %v", want)
	}
}
