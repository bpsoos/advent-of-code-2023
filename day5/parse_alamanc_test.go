package day5_test

import (
	"advent2023/day5"
	"reflect"
	"strings"
	"testing"
)

func TestParseAlmanacWithSingleMap(t *testing.T) {
	want := &day5.Almanac{
		Seeds: []int{1, 2},
		Maps:  []day5.AlmanacMap{{day5.AlmanacMapEntry{1, 2, 3}, day5.AlmanacMapEntry{3, 2, 1}}},
	}
	rawAlmanac := strings.TrimSpace(`
seeds: 1 2

seed-to-soil map:
1 2 3
3 2 1
`)
	almanac := day5.ParseAlmanac(strings.NewReader(rawAlmanac))
	if !reflect.DeepEqual(almanac, want) {
		t.Errorf(
			"parsed almanac doesn't match expected\nparsed:  %v\nexpected: %v",
			almanac,
			want,
		)
	}
}

func TestParseAlmanacWithMultipleMaps(t *testing.T) {
	want := &day5.Almanac{
		Seeds: []int{34, 52, 22},
		Maps: []day5.AlmanacMap{
			{
				day5.AlmanacMapEntry{1, 2, 3},
				day5.AlmanacMapEntry{3, 2, 1},
			},
			{
				day5.AlmanacMapEntry{12, 23, 0},
				day5.AlmanacMapEntry{1, 99, 77},
			},
		},
	}
	rawAlmanac := strings.TrimSpace(`
seeds: 34 52 22

seed-to-soil map:
1 2 3
3 2 1

soil-to-fertilizer map:
12 23 0
1 99 77
`)
	almanac := day5.ParseAlmanac(strings.NewReader(rawAlmanac))
	if !reflect.DeepEqual(almanac, want) {
		t.Errorf(
			"parsed almanac doesn't match expected\nparsed:  %v\nexpected: %v",
			almanac,
			want,
		)
	}
}
