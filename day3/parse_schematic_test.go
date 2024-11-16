package day3_test

import (
	"advent2023/day3"
	"advent2023/mathx"
	"reflect"
	"strings"
	"testing"
)

func TestParseSimpleSchematic(t *testing.T) {
	rawSchematic := strings.TrimSpace(`
467..
...*.
..35.
`)
	want := day3.NewSchematic()
	want.AddSymbol(mathx.Point{
		X: 3,
		Y: 1,
	})

	want.PartNumbers = []day3.PartNumber{
		{
			Value: 467,
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
				Start: mathx.Point{
					X: 2,
					Y: 2,
				}, End: mathx.Point{
					X: 3,
					Y: 2,
				}},
		},
	}
	schematic := day3.ParseSchematic(strings.NewReader(rawSchematic))
	if !reflect.DeepEqual(schematic, want) {
		t.Errorf("parsed schematic does not match expected")
		t.Logf("parsed: %v", schematic)
		t.Logf("expected: %v", want)
	}
}

func TestParseComplexSchematic(t *testing.T) {
	rawSchematic := strings.TrimSpace(`
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
`)
	want := day3.NewSchematic()
	want.AddSymbol(mathx.Point{
		X: 3,
		Y: 1,
	})
	want.AddSymbol(mathx.Point{
		X: 6,
		Y: 3,
	})
	want.AddSymbol(mathx.Point{
		X: 3,
		Y: 4,
	})
	want.AddSymbol(mathx.Point{
		X: 5,
		Y: 5,
	})
	want.AddSymbol(mathx.Point{
		X: 3,
		Y: 8,
	})
	want.AddSymbol(mathx.Point{
		X: 5,
		Y: 8,
	})

	want.PartNumbers = []day3.PartNumber{
		{
			Value: 467,
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
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
			Position: day3.PartNumberPosition{
				Start: mathx.Point{
					X: 5,
					Y: 9,
				}, End: mathx.Point{
					X: 7,
					Y: 9,
				}},
		},
	}
	schematic := day3.ParseSchematic(strings.NewReader(rawSchematic))
	if !reflect.DeepEqual(schematic, want) {
		t.Errorf("parsed schematic does not match expected")
		t.Logf("parsed: %v", schematic)
		t.Logf("expected: %v", want)
	}
}
