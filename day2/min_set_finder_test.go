package day2_test

import (
	"advent2023/day2"
	"reflect"
	"testing"
)

func TestFindMinSet(t *testing.T) {
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
	want := day2.MinSets{
		{
			Red:   4,
			Green: 2,
			Blue:  6,
		},
		{
			Red:   1,
			Green: 3,
			Blue:  4,
		},
	}
	minSets := day2.MinSetFinder{}.Find(games)
	if !reflect.DeepEqual(want, minSets) {
		t.Errorf("minsets doesn't match expected")
		t.Logf("minsets: %v", minSets)
		t.Logf("expected: %v", want)
	}
}
