package day3_test

import (
	"advent2023/day3"
	"advent2023/mathx"
	"reflect"
	"testing"
)

func TestFindPartNumbersWithAllNumbersMatching(t *testing.T) {
	/*
	   467..
	   ...*.
	   ..35.
	*/
	schematic := day3.NewSchematic()
	schematic.AddSymbol(mathx.Point{
		X: 3,
		Y: 1,
	})

	schematic.PartNumbers = []day3.PartNumber{
		{
			Value: 467,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 0,
					Y: 0,
				}, End: mathx.Point{
					X: 2,
					Y: 0,
				}},
		},
		{
			Value: 35,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 2,
					Y: 2,
				}, End: mathx.Point{
					X: 3,
					Y: 2,
				}},
		},
	}
	want := []int{467, 35}
	numbers := day3.PartNumberFinder{}.Find(schematic)
	if !reflect.DeepEqual(want, numbers) {
		t.Error("part numbers don't match expected")
		t.Logf("part numbers: %v", numbers)
		t.Logf("expected: %v", want)
	}
}

func TestFindPartNumbersWithNoNumbersMatching(t *testing.T) {
	/*
	   467..
	   .....
	   ..35.
	*/
	schematic := day3.NewSchematic()

	schematic.PartNumbers = []day3.PartNumber{
		{
			Value: 467,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 0,
					Y: 0,
				}, End: mathx.Point{
					X: 2,
					Y: 0,
				}},
		},
		{
			Value: 35,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 2,
					Y: 2,
				}, End: mathx.Point{
					X: 3,
					Y: 2,
				}},
		},
	}
	want := make([]int, 0, 3)
	numbers := day3.PartNumberFinder{}.Find(schematic)
	if !reflect.DeepEqual(want, numbers) {
		t.Error("part numbers don't match expected")
		t.Logf("part numbers: %v", numbers)
		t.Logf("expected: %v", want)
	}
}

func TestFindPartNumbersWithComplexSchematic(t *testing.T) {
	/*
	   467..114..
	   ...*......
	   ..35..633.
	   ......#...
	   617*......
	   .....+.58.
	   ..592.....
	   ......755.
	   ...$.*....
	   .664.598..
	*/
	schematic := day3.NewSchematic()
	schematic.AddSymbol(mathx.Point{
		X: 3,
		Y: 1,
	})
	schematic.AddSymbol(mathx.Point{
		X: 6,
		Y: 3,
	})
	schematic.AddSymbol(mathx.Point{
		X: 3,
		Y: 4,
	})
	schematic.AddSymbol(mathx.Point{
		X: 5,
		Y: 5,
	})
	schematic.AddSymbol(mathx.Point{
		X: 3,
		Y: 8,
	})
	schematic.AddSymbol(mathx.Point{
		X: 5,
		Y: 8,
	})

	schematic.PartNumbers = []day3.PartNumber{
		{
			Value: 467,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 0,
					Y: 0,
				}, End: mathx.Point{
					X: 2,
					Y: 0,
				}},
		},
		{
			Value: 114,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 5,
					Y: 0,
				}, End: mathx.Point{
					X: 7,
					Y: 0,
				}},
		},
		{
			Value: 35,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 2,
					Y: 2,
				}, End: mathx.Point{
					X: 3,
					Y: 2,
				}},
		},
		{
			Value: 633,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 6,
					Y: 2,
				}, End: mathx.Point{
					X: 8,
					Y: 2,
				}},
		},
		{
			Value: 617,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 0,
					Y: 4,
				}, End: mathx.Point{
					X: 2,
					Y: 4,
				}},
		},
		{
			Value: 58,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 7,
					Y: 5,
				}, End: mathx.Point{
					X: 8,
					Y: 5,
				}},
		},
		{
			Value: 592,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 2,
					Y: 6,
				}, End: mathx.Point{
					X: 4,
					Y: 6,
				}},
		},
		{
			Value: 755,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 6,
					Y: 7,
				}, End: mathx.Point{
					X: 8,
					Y: 7,
				}},
		},
		{
			Value: 664,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 1,
					Y: 9,
				}, End: mathx.Point{
					X: 3,
					Y: 9,
				}},
		},
		{
			Value: 598,
			Position: day3.PartPosition{
				Start: mathx.Point{
					X: 5,
					Y: 9,
				}, End: mathx.Point{
					X: 7,
					Y: 9,
				}},
		},
	}
	want := []int{467, 35, 633, 617, 592, 755, 664, 598}
	numbers := day3.PartNumberFinder{}.Find(schematic)
	if !reflect.DeepEqual(want, numbers) {
		t.Error("part numbers don't match expected")
		t.Logf("part numbers: %v", numbers)
		t.Logf("expected: %v", want)
	}
}
