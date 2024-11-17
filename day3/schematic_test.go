package day3_test

import (
	"advent2023/day3"
	"advent2023/mathx"
	"reflect"
	"testing"
)

func TestAdjPositions(t *testing.T) {
	partNumber := day3.PartNumber{
		Value: 35,
		Position: day3.PartPosition{
			Start: mathx.Point{
				X: 2,
				Y: 2,
			},
			End: mathx.Point{
				X: 3,
				Y: 2,
			},
		},
	}
	want := []mathx.Point{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},

		{X: 4, Y: 1},
		{X: 4, Y: 2},
		{X: 4, Y: 3},

		{X: 2, Y: 1},
		{X: 2, Y: 3},

		{X: 3, Y: 1},
		{X: 3, Y: 3},
	}
	adjPositions := partNumber.AdjPositions()
	if !reflect.DeepEqual(want, adjPositions) {
		t.Error("adj positions don't match expected")
		t.Logf("adj positions: %v", adjPositions)
		t.Logf("expected: %v", want)
	}
}
